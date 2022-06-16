// Package product holds the repository and the implementations for a Repository
package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hollson/ddd1/domain/agg"
)

var (
	ErrProductNotFound     = errors.New("the product was not found")
	ErrProductAlreadyExist = errors.New("the product already exists")
)

// Repository 产品「仓储接口」
type Repository interface {
	GetAll() ([]agg.Product, error)
	GetByID(id uuid.UUID) (agg.Product, error)
	Add(product agg.Product) error
	Update(product agg.Product) error
	Delete(id uuid.UUID) error
}
