package agg

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hollson/ddd1/domain/entity"

	"github.com/hollson/ddd1/vo"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid Person")
)

// Customer 「聚合实体」客户
type Customer struct {
	Person       *entity.Person   `bson:"Person"`
	Products     []*entity.Item   `bson:"Products"`     // 一个客户可以持有许多产品
	Transactions []vo.Transaction `bson:"Transactions"` // 一个客户可以执行许多事务
}

// NewCustomer 是创建新的Customer聚合的工厂
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// 创建一个新person并生成ID
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}
	// 创建一个customer对象并初始化所有的值以避免空指针异常
	return Customer{
		Person:       person,
		Products:     make([]*entity.Item, 0),
		Transactions: make([]vo.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.Person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.Person == nil {
		c.Person = &entity.Person{}
	}
	c.Person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.Person == nil {
		c.Person = &entity.Person{}
	}
	c.Person.Name = name
}

func (c *Customer) GetName() string {
	return c.Person.Name
}
