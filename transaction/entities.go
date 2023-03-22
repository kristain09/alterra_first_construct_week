package transaction

import "time"

type Transactions struct {
	ID          int
	Invoice     int
	TransDate   time.Time
	Total       int
	CustomersID int
	CreatedBy   int
}
