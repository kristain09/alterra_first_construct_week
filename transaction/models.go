package transaction

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type TransactionsModels struct {
	conn *sql.DB
}

func (tm *TransactionsModels) SetConnTransModels(db *sql.DB) {
	tm.conn = db
}

func (tm TransactionsModels) PrintAllTransData() error {
	var transactions []Transactions

	query := "SELECT id, invoice, transdate, total, customers_id, created_by FROM transactions WHERE deleted_at IS NULL"
	rows, err := tm.conn.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transactions
		err := rows.Scan(&transaction.ID, &transaction.Invoice, &transaction.TransDate, &transaction.Total, &transaction.CustomersID, &transaction.CreatedBy)

		if err != nil {
			return err
		}

		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Invoice", "Transdate", "Total", "Customers ID", "Created By"})

	for _, transaction := range transactions {
		row := []string{
			strconv.Itoa(transaction.ID),
			strconv.Itoa(transaction.Invoice),
			transaction.TransDate.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%d", transaction.Total),
			strconv.Itoa(transaction.CustomersID),
			strconv.Itoa(transaction.CreatedBy),
			"-",
		}
		table.Append(row)
	}

	table.Render()

	return nil
}

func (tm TransactionsModels) PrintTransDataByUserID(id int) error {
	var transactions []Transactions

	query := "SELECT id, invoice, transdate, total, customers_id, created_by, deleted_at FROM transactions WHERE created_by = ? AND deleted_at IS NULL"
	rows, err := tm.conn.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transactions
		err := rows.Scan(&transaction.ID, &transaction.Invoice, &transaction.TransDate, &transaction.Total, &transaction.CustomersID, &transaction.CreatedBy)
		if err != nil {
			return err
		}
		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Invoice", "Transdate", "Total", "Customers ID", "Created By", "Deleted At"})

	for _, transaction := range transactions {
		row := []string{
			strconv.Itoa(transaction.ID),
			strconv.Itoa(transaction.Invoice),
			transaction.TransDate.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%d", transaction.Total),
			strconv.Itoa(transaction.CustomersID),
			strconv.Itoa(transaction.CreatedBy),
			"-",
		}
		table.Append(row)
	}

	table.Render()

	return nil
}

func (tm *TransactionsModels) InitDeletedAt(id int) error {

	_, err := tm.conn.Exec("UPDATE transactions SET deleted_at = current_timestamp  WHERE id = ?", id)

	if err != nil {
		return err
	}
	return nil
}
