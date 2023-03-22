package transactions

import (
	"fmt"
	"time"
)

type TransactionController struct {
	transactionModel *TransactionsModel
}

func NewTransactionController(tm *TransactionsModel) *TransactionController {
	return &TransactionController{tm}
}

func (tc *TransactionController) HandleRequest() {
	now := time.Now()
	fmt.Println(" Product Information\n", now.Format("Monday, 2006 January 2 15:04:05"))
	tc.handleCreateTransaction()
}

func (tc *TransactionController) handleCreateTransaction() {
	var productID, customersID, createdBy, quantity, total int

	fmt.Println("---------------------------")
	fmt.Println("Create Transaction")
	fmt.Println("---------------------------")

	fmt.Print("Enter product ID: ")
	fmt.Scan(&productID)

	fmt.Print("Enter customer ID: ")
	fmt.Scan(&customersID)

	fmt.Print("Enter created by user ID: ")
	fmt.Scan(&createdBy)

	fmt.Print("Enter quantity: ")
	fmt.Scan(&quantity)

	fmt.Print("Enter total: ")
	fmt.Scan(&total)

	transaction, err := tc.transactionModel.CreateTransaction(productID, customersID, createdBy, quantity, total)
	if err != nil {
		fmt.Println("Failed to create transaction:", err)
		return
	}

	fmt.Println("---------------------------")
	fmt.Printf("Transaction created successfully with details:\nInvoice\t: %s\nTransdate\t: %s\nProduct ID\t: %d\nCustomer ID\t: %d\nQuantity\t: %d\nTotal\t\t: %d\nCreated by\t: %d\n", transaction.Invoice, transaction.Transdate, transaction.ProductID, transaction.CustomersID, transaction.Quantity, transaction.Total, transaction.CreatedBy)
	fmt.Println("---------------------------")
}
