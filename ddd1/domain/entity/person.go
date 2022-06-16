package entity

import (
	"fmt"

	"github.com/google/uuid"
)

// Person 「领域实体/根实体」表示领域中的人
type Person struct {
	ID   uuid.UUID `json:"id" bson:"id"`     // 实体标识符
	Name string    `json:"name" bson:"name"` // 名字
	Age  int       `json:"age" name:"age"`   // 年龄
}

func (p *Person) String() string {
	return fmt.Sprintf("%v", *p)
}
