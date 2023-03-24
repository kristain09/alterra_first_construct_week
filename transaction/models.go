package transaction

import (
	"database/sql"
	"log"
	"time"
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
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transactions
		err := rows.Scan(&transaction.ID, &transaction.Invoice, &transaction.TransDate, &transaction.Total, &transaction.CustomersID, &transaction.CreatedBy)
		if err != nil {
			return nil, err
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

func (tm TransactionsModels) InsertTransaction(input Transactions, userID int) (int, error) {
	tx, err := tm.db.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	stmt, err := tx.Prepare("INSERT INTO transactions(invoice, transdate, total, customers_id, created_by) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(input.Invoice, time.Now(), input.Total, input.CustomersID, userID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return 0, err
	}

	stmt2, err := tx.Prepare("INSERT INTO transaction_product(transaction_id, product_id, qty) VALUES (?,?,?)")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return 0, err
	}

	for _, product := range input.Product {
		_, err = stmt2.Exec(lastInsertID, product.ID, product.Stock)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return 0, err
		}

		_, err = tx.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", product.Stock, product.ID)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return 0, err
	}

	return int(lastInsertID), nil
}
