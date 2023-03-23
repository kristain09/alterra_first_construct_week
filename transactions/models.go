package transactions

import (
	"database/sql"
	"fmt"
)

type TransactionsModel struct {
	conn *sql.DB
}

func (tm *TransactionsModel) SetConnection(db *sql.DB) {
	tm.conn = db
}

func (tm *TransactionsModel) CreateTransaction(product_id int, customers_id int, created_by int, quantity int, total int) (Transactions, error) {

	query := "INSERT INTO transactions (invoice, transdate, product_id, customers_id, created_by, quantity, total) VALUES (generate_invoice_number(), NOW(), ?, ?, ?, ?, ?)"
	result, err := tm.conn.Exec(query, product_id, customers_id, created_by, quantity, total)
	if err != nil {
		return Transactions{}, fmt.Errorf("error creating transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Transactions{}, fmt.Errorf("error getting transaction ID: %w", err)
	}

	query = "SELECT id, invoice, transdate, product_id, customers_id, created_by, quantity, total FROM transactions WHERE id = ?"
	row := tm.conn.QueryRow(query, id)
	transaction := Transactions{}
	err = row.Scan(&transaction.ID, &transaction.Invoice, &transaction.Transdate, &transaction.ProductID, &transaction.CustomersID, &transaction.CreatedBy, &transaction.Quantity, &transaction.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return Transactions{}, fmt.Errorf("transactions with ID %d not found", id)
		}
		return Transactions{}, fmt.Errorf("error getting transaction details: %w", err)
	}
	return transaction, nil
}
