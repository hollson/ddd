package services

import (
	"fmt"

	"github.com/google/uuid"
)

type Tavern struct {
	OrderService *OrderService
	BillingService interface{}
}

func NewTavern(opts ...TavernOption) (*Tavern, error) {
	t := &Tavern{}
	for _, opt := range opts {
		err := opt(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// Order 客户执行订单
func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products...)
	if err != nil {
		return err
	}
	fmt.Printf("向客户收费: %0.0f\n", price)
	return nil
}
