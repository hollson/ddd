package aggregate

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/entity"
)

var (
	ErrMissingValues = errors.New("missing values")
)

// Product 聚合了价格和数量
type Product struct {
	item     *entity.Item
	price    float64 // 价格
	quantity int     // 库存
}

// NewProduct will create a new product
// will return error if name of description is empty
func NewProduct(name, description string, price float64) *Product {
	if name == "" || description == "" {
		panic(ErrMissingValues)
	}
	return &Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p *Product) String() string {
	return fmt.Sprintf("%v,%v,%v", p.item, p.price, p.quantity)
}
