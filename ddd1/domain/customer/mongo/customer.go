// Mongo is a mongo implementation of the Customer CustomerRepo
package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hollson/ddd1/domain/agg"
)

type CustomerRepo struct {
	// db *mongo.Database
	// customer *mongo.Collection
}

type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c agg.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// ToAggregate converts into a aggregate.Customer
// this could validate all values present etc
func (m mongoCustomer) ToAggregate() agg.Customer {
	c := agg.Customer{}
	c.SetID(m.ID)
	c.SetName(m.Name)
	return c
}

func New(ctx context.Context, conn string) (*CustomerRepo, error) {
	panic("implement")
}

func (mr *CustomerRepo) Get(id uuid.UUID) (agg.Customer, error) {
	panic("implement")
}

func (mr *CustomerRepo) Add(c agg.Customer) error {
	panic("implement")
}

func (mr *CustomerRepo) Update(c agg.Customer) error {
	panic("implement")
}
