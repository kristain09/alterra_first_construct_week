package transactions

type Transactions struct {
	ID          int
	Invoice     string
	Transdate   string
	Quantity    int
	Total       int
	ProductID   int
	CustomersID int
	CreatedBy   int
}
