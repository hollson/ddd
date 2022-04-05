// Package memory is a in-memory implementation of the customer repository
package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
	"github.com/percybolmer/ddd-go/domain/customer"
)

// CustomerRepo fulfills the CustomerRepository interface
type CustomerRepo struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *CustomerRepo {
	return &CustomerRepo{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get finds a customer by ID
func (mr *CustomerRepo) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

// Add will add a new customer to the repository
func (mr *CustomerRepo) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		// Saftey check if customers is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *CustomerRepo) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
