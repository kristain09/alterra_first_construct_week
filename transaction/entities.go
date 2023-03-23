package transaction

import (
	"first_construct_week/products"
	"time"
)

type Transactions struct {
	ID          int
	Invoice     int
	TransDate   time.Time
	Total       int
	CustomersID int
	CreatedBy   int
	Product     []products.Products
}
