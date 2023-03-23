package transactions

import (
	"database/sql"
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
