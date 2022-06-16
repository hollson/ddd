// Package services 包含将存储库连接到业务流中的所有服务
package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hollson/ddd1/domain/agg"
	"github.com/hollson/ddd1/domain/customer"
	"github.com/hollson/ddd1/domain/product"
)

// OrderService 聚合关系
type OrderService struct {
	Customers customer.Repository
	Products  product.Repository
}

func NewOrderService(cfgs ...OrderOption) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs ...uuid.UUID) (float64, error) {
	c, err := o.Customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []agg.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.Products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	fmt.Printf("顾客: %s「%s」 下单%d件商品，合计消费%v 元\n", c.GetName(), c.GetID(), len(products), price)

	// Add Products and Update Customer
	return price, nil
}
