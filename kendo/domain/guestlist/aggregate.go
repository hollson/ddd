package guestlist

import (
	"context"
	"fmt"
	"time"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/uuid"
)

func init() {
	eh.RegisterAggregate(func(id uuid.UUID) eh.Aggregate {
		return NewInvitationAggregate(id)
	})
}

// InvitationAggregateType is the type name of the aggregate.
const InvitationAggregateType eh.AggregateType = "Invitation"

// InvitationAggregate 聚合根.
// 聚合根将保护邀请只能被接受或拒绝，但不能同时接受或拒绝。
type InvitationAggregate struct {
	*events.AggregateBase
	name string // 姓名
	age  int    // 年龄

	// TODO: Replace with FSM.
	accepted  bool // 接受
	declined  bool // 谢绝
	confirmed bool // 确认
	denied    bool // 拒绝
}

// 校验，保证Invitation实现了Aggregate接口
var _ = eh.Aggregate(&InvitationAggregate{})

// NewInvitationAggregate 创建Invitation聚合根对象
func NewInvitationAggregate(id uuid.UUID) *InvitationAggregate {
	return &InvitationAggregate{AggregateBase: events.NewAggregateBase(InvitationAggregateType, id)}
}

// HandleCommand 将自定义事件追加(注册)到聚合根到事件列表中
func (a *InvitationAggregate) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd := cmd.(type) {
	case *CreateInvite:
		a.AppendEvent(InviteCreatedEvent, &InviteCreatedData{cmd.Name, cmd.Age}, time.Now())
		return nil
	case *AcceptInvite:
		if a.name == "" {
			return fmt.Errorf("invitee does not exist")
		}
		if a.declined {
			return fmt.Errorf("%s already declined", a.name)
		}
		if a.accepted {
			return nil
		}
		a.AppendEvent(InviteAcceptedEvent, nil, time.Now())
		return nil
	case *DeclineInvite:
		if a.name == "" {
			return fmt.Errorf("invitee does not exist")
		}

		if a.accepted {
			return fmt.Errorf("%s already accepted", a.name)
		}

		if a.declined {
			return nil
		}

		a.AppendEvent(InviteDeclinedEvent, nil, time.Now())
		return nil

	case *ConfirmInvite:
		if a.name == "" {
			return fmt.Errorf("invitee does not exist")
		}

		if !a.accepted || a.declined {
			return fmt.Errorf("only accepted invites can be confirmed")
		}

		a.AppendEvent(InviteConfirmedEvent, nil, time.Now())
		return nil

	case *DenyInvite:
		if a.name == "" {
			return fmt.Errorf("invitee does not exist")
		}

		if !a.accepted || a.declined {
			return fmt.Errorf("only accepted invites can be denied")
		}

		a.AppendEvent(InviteDeniedEvent, nil, time.Now())
		return nil
	}
	return fmt.Errorf("couldn't handle command")
}

// ApplyEvent 通过设置其值将事件应用于聚合，如果没有错误，则应通过调用IncrementVersion来增加版本。
func (a *InvitationAggregate) ApplyEvent(ctx context.Context, event eh.Event) error {
	switch event.EventType() {
	case InviteCreatedEvent:
		if data, ok := event.Data().(*InviteCreatedData); ok {
			a.name = data.Name
			a.age = data.Age
		} else {
			// log.Println("invalid event data type:", event.Data())
		}
	case InviteAcceptedEvent:
		a.accepted = true
	case InviteDeclinedEvent:
		a.declined = true
	case InviteConfirmedEvent:
		a.confirmed = true
	case InviteDeniedEvent:
		a.denied = true
	}
	return nil
}
