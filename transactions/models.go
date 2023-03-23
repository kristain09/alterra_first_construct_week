package transactions

import (
	"database/sql"
	"fmt"
)

type TransactionsModels struct {
	conn *sql.DB
}

func (tm *TransactionsModels) SetConnDBTransModels(db *sql.DB) {
	tm.conn = db
}

func (tm TransactionsModels) GetAllTransData() ([]Transactions, error) {
	var transactions []Transactions

	query := "SELECT id, invoice, transdate, total, customers_id, created_by FROM transactions WHERE deleted_at IS NULL"
	rows, err := tm.conn.Query(query)
	if err != nil {
		return transactions, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transactions
		err := rows.Scan(&transaction.ID, &transaction.Invoice, &transaction.TransDate, &transaction.Total, &transaction.CustomersID, &transaction.CreatedBy)
		if err != nil {
			return transactions, err
		}

		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (tm TransactionsModels) GetAllTransDataByID(id int) ([]Transactions, error) {
	var transactions []Transactions

	query := "SELECT id, invoice, transdate, total, customers_id, created_by, deleted_at FROM transactions WHERE created_by = ? AND deleted_at IS NULL"
	rows, err := tm.conn.Query(query, id)
	if err != nil {
		return transactions, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transactions
		err := rows.Scan(&transaction.ID, &transaction.Invoice, &transaction.TransDate, &transaction.Total, &transaction.CustomersID, &transaction.CreatedBy)
		if err != nil {
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (tm *TransactionsModels) InitDeletedAt(id int) error {

	_, err := tm.conn.Exec("UPDATE transactions SET deleted_at = current_timestamp  WHERE id = ?", id)

	if err != nil {
		return err
	}
	return nil
}

func (tm *TransactionsModels) CreateTransaction(product_id int, customers_id int, created_by int, quantity int, total int) (Transactions, error) {

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
	err = row.Scan(&transaction.ID, &transaction.Invoice, &transaction.TransDate, &transaction.Product, &transaction.CustomersID, &transaction.CreatedBy, &transaction.Quantity, &transaction.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return Transactions{}, fmt.Errorf("transactions with ID %d not found", id)
		}
		return Transactions{}, fmt.Errorf("error getting transaction details: %w", err)
	}
	return transaction, nil
}
