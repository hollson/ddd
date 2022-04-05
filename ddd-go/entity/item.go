package entity

import (
	"fmt"

	"github.com/google/uuid"
)

// Item 「领域实体/根实体」」表示子领域的根实体
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}

func (p *Item) String() string {
	return fmt.Sprintf("%v", *p)
}
