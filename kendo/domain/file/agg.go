package file

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hollson/kendo/infrastructure/bus"
	"github.com/hollson/kendo/infrastructure/repo"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/sirupsen/logrus"
)

const (
	AgentAggregateType = eh.AggregateType("AggregateType_Files")
)

// Agg 「文件领域」聚合对象
type Agg struct {
	*events.AggregateBase
}

var SingleFilesAgg *Agg

func init() {
	// 创建/注册聚合
	agg := &Agg{AggregateBase: events.NewAggregateBase(AgentAggregateType, uuid.Nil)}
	// 向聚合根添加事件
	bus.RegisterHandler(AddFileCmdType, agg)

	SingleFilesAgg = agg
}

// Command异步执行，不需要返回值的
func (a *Agg) HandleCommand(ctx context.Context, cmd eh.Command) (err error) {
	switch cmd := cmd.(type) {
	case *AddFileCmd:
		vo := FileValue{
			FileInfo: repo.FileInfo{
				Id:          uuid.New().String(),
				FileName:    cmd.FileName,
				Size:        cmd.Size,
				ContentType: cmd.ContentType,
			},
		}
		en := newEntity(vo)
		err = en.AddFile(cmd.FileBody)
		if err != nil {
			logrus.Errorf("新增文件出错：%v ", err.Error())
		}
	default:
		err = fmt.Errorf("couldn't handle command")
	}
	return
}

func (a *Agg) ApplyEvent(ctx context.Context, event eh.Event) (err error) {
	return
}

// Command同步执行，需要返回值的
func (a *Agg) DealCommand(ctx context.Context, cmd eh.Command) (interface{}, error) {
	return nil, fmt.Errorf("couldn't Dealer command")
}

// 聚合根对外开放的能力
func (a *Agg) AddNewFile(vo FileValue) (fileId string, err error) {
	fileId = uuid.New().String()
	vo.Id = fileId
	en := newEntity(vo)
	err = en.AddFile(vo.FileBody)
	return
}
