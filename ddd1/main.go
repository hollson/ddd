// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hollson/ddd1/domain/agg"

	"github.com/hollson/ddd1/services"
)

func makeProducts() []agg.Product {
	p1 := agg.NewProduct("茅台", "这是白酒", 199)
	p2 := agg.NewProduct("德芙", "这是巧克力", 0.99)
	p3 := agg.NewProduct("乐事", "这是薯片", 2.88)
	products := []agg.Product{*p1, *p2, *p3}
	return products
}

func ordertest() {
	customer, _ := agg.NewCustomer("Percy") // 创建顾客
	products := makeProducts()              // 创建商品

	// 订单服务,将顾客和商品关联起来
	_orderService, err := services.NewOrderService(
		services.WithMemoryCustomerRepository(),
		services.WithMemoryProductRepository(products...),
	)

	// 把顾客信息添加到订单服务中
	err = _orderService.Customers.Add(customer)
	if err != nil {
		// t.Error(err)
	}

	// Perform Order for one beer
	order := []uuid.UUID{products[0].GetID(), products[1].GetID(), products[2].GetID()}

	orderDetails, err := _orderService.CreateOrder(customer.GetID(), order...)
	if err != nil {
	}
	fmt.Println(orderDetails)
}

func tavernt1() {
	products := makeProducts()
	os, err := services.NewOrderService(
		services.WithMemoryCustomerRepository(),
		// WithMongoCustomerRepository("mongodb://localhost:27017"),
		services.WithMemoryProductRepository(products...),
	)

	tavern, err := services.NewTavern(services.WithOrderService(os))
	if err != nil {
		// t.Error(err)
	}

	cust, err := agg.NewCustomer("古灵精怪")
	if err != nil {
	}

	err = os.Customers.Add(cust)
	if err != nil {
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		// t.Error(err)
	}
}

func tavern2() {
	// Create OrderService
	products := makeProducts()

	os, err := services.NewOrderService(
		services.WithMemoryCustomerRepository(),
		services.WithMemoryProductRepository(products...),
	)
	if err != nil {
		// t.Error(err)
	}

	tavern, err := services.NewTavern(services.WithOrderService(os))
	if err != nil {
		// t.Error(err)
	}

	cust, err := agg.NewCustomer("Percy")
	if err != nil {
		// t.Error(err)
	}

	err = os.Customers.Add(cust)
	if err != nil {
		// t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		// t.Error(err)
	}
}

func main() {
	ordertest()
	tavernt1()
	tavern2()
}
