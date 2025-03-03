package file

import (
	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
)

const (
	AddFileCmdType eh.CommandType = "AddFileCmd"
)

func init() {
	eh.RegisterCommand(func() eh.Command { return &AddFileCmd{} })
}

// 添加文件
type AddFileCmd struct {
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
	Size        int    `json:"size"`
	FileBody    []byte `json:"file_body" eh:"optional"`
}

func (c *AddFileCmd) AggregateID() uuid.UUID {
	return uuid.Nil
}
func (c *AddFileCmd) CommandType() eh.CommandType     { return AddFileCmdType }
func (c *AddFileCmd) AggregateType() eh.AggregateType { return "" }
func (c *AddFileCmd) Verify() error                   { return nil }
