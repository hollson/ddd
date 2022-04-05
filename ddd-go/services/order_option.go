// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package services

import (
	"context"

	"github.com/percybolmer/ddd-go/aggregate"
	"github.com/percybolmer/ddd-go/domain/customer"
	"github.com/percybolmer/ddd-go/domain/customer/memory"
	"github.com/percybolmer/ddd-go/domain/customer/mongo"
	prodmemory "github.com/percybolmer/ddd-go/domain/product/memory"
)

type OrderOption func(os *OrderService) error

func WithCustomerRepository(cr customer.Repository) OrderOption {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.Customers = cr
		return nil
	}
}

func WithMongoCustomerRepository(connectionString string) OrderOption {
	return func(os *OrderService) error {
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.Customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderOption {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products ...aggregate.Product) OrderOption {
	return func(os *OrderService) error {
		pr := prodmemory.New()

		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.Products = pr
		return nil
	}
}