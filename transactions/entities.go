package transactions

import "database/sql"

type Transactions struct {
	ID          int
	Invoice     sql.NullString
	TransDate   string
	Total       int
	CustomersID int
	CreatedBy   int
	Product     int
	Quantity    int
}
