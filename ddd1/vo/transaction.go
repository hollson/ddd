package vo

import (
	"time"

	"github.com/google/uuid"
)

// Transaction 「值对象」
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
