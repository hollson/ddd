// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mongo

import (
	"github.com/google/uuid"
)

type ProductRepo struct{}

func (p *ProductRepo) GetAll() ([]agg.Product, error) {
	panic("implement me")
}

func (p *ProductRepo) GetByID(id uuid.UUID) (agg.Product, error) {
	panic("implement me")
}

func (p *ProductRepo) Add(product agg.Product) error {
	panic("implement me")
}

func (p *ProductRepo) Update(product agg.Product) error {
	panic("implement me")
}

func (p *ProductRepo) Delete(id uuid.UUID) error {
	panic("implement me")
}
