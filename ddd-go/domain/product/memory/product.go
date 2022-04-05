// Package memory is a in memory implementation of the ProductRepo interface.
package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
	"github.com/percybolmer/ddd-go/domain/product"
)

type ProductRepo struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *ProductRepo {
	return &ProductRepo{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

// GetAll returns all products as a slice
// Yes, it never returns an error, but
// A database implementation could return an error for instance
func (mpr *ProductRepo) GetAll() ([]aggregate.Product, error) {
	// Collect all Products from map
	var products []aggregate.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

// GetByID searches for a product based on it's ID
func (mpr *ProductRepo) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

// Add will add a new product to the repository
func (mpr *ProductRepo) Add(newprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	mpr.products[newprod.GetID()] = newprod

	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *ProductRepo) Update(upprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

// Delete remove an product from the repository
func (mpr *ProductRepo) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
