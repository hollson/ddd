// Package product holds the repository and the implementations for a Repository
package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
)

var (
	ErrProductNotFound     = errors.New("the product was not found")
	ErrProductAlreadyExist = errors.New("the product already exists")
)

// Repository 产品「仓储接口」
type Repository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
