package file

import (
	"context"
	"fmt"


	"github.com/hollson/kendo/infrastructure/bus"
	"github.com/hollson/kendo/infrastructure/repo"
	eh "github.com/looplab/eventhorizon"
)

func init() {
	// 注册事件handler
	bus.RegisterEventHandler(eh.MatchEvents{FileAddedEvent}, &filesEventHandler{})

}

type filesEventHandler struct {
}

func (a *filesEventHandler) HandlerType() eh.EventHandlerType {
	return ""
}
func (a *filesEventHandler) HandleEvent(ctx context.Context, event eh.Event) (err error) {
	switch evtData := event.Data().(type) {
	case *repo.FileInfo:
		println("订阅者发现了，文件新增成功！ 文件ID：", evtData.Id, "订阅者继续业务逻辑...")
	}

	return fmt.Errorf("couldn't handle Event")
}
