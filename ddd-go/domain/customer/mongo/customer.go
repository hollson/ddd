// Mongo is a mongo implementation of the Customer CustomerRepo
package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
)

type CustomerRepo struct {
	// db *mongo.Database
	// customer *mongo.Collection
}

type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// ToAggregate converts into a aggregate.Customer
// this could validate all values present etc
func (m mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}
	c.SetID(m.ID)
	c.SetName(m.Name)
	return c
}

func New(ctx context.Context, conn string) (*CustomerRepo, error) {
	panic("implement")
}

func (mr *CustomerRepo) Get(id uuid.UUID) (aggregate.Customer, error) {
	panic("implement")
}

func (mr *CustomerRepo) Add(c aggregate.Customer) error {
	panic("implement")
}

func (mr *CustomerRepo) Update(c aggregate.Customer) error {
	panic("implement")
}
