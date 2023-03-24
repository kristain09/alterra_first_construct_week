package transaction

import (
	"first_construct_week/products"
)

type Transactions struct {
	ID          int
	Invoice     int
	TransDate   string
	Total       int
	CustomersID int
	CreatedBy   int
	Product     []products.Products
}
