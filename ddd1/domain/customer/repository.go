package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hollson/ddd1/domain/agg"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
)

// Repository 客户「仓储接口」
type Repository interface {
	Get(uuid.UUID) (agg.Customer, error)
	Add(agg.Customer) error
	Update(agg.Customer) error
}
