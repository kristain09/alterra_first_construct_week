package transactions

import (
	"fmt"
	"log"
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
	log.Println(" Transaction Information\n", now.Format("Monday, 2006 January 2 15:04:05"))
	tc.handleCreateTransaction()
}

func (tc *TransactionController) handleCreateTransaction() {
	var productID, customersID, createdBy, quantity, total int

	log.Println("---------------------------")
	log.Println("Create Transaction")
	log.Println("---------------------------")

	for {
		log.Print("Add another product? (y/n): ")
		var choice string
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}
		if choice == "n" {
			break
		}

		log.Print("Enter product ID: ")
		if _, err := fmt.Scan(&productID); err != nil {
			log.Println("Invalid input, please try again.")
			continue
		}

		log.Print("Enter quantity: ")
		if _, err := fmt.Scan(&quantity); err != nil {
			fmt.Println("Invalid input, please try again.")
			continue

			log.Print("Enter total: ")
			if _, err := fmt.Scan(&total); err != nil {
				log.Println("Invalid input, please try again.")
				continue
			}
		}

		log.Print("Enter customer ID: ")
		fmt.Scan(&customersID)

		log.Print("Enter created by user ID: ")
		fmt.Scan(&createdBy)

		// // Query database disini
		// product, err := tc.transactionModel.GetProduct(productID)
		// if err != nil {
		// 	log.Println("Failed to get product:", err)
		// 	return
		// }

		// // Calculate the total by multiplying the price by the quantity
		// total := product.Price * quantity

		transaction, err := tc.transactionModel.CreateTransaction(productID, customersID, createdBy, quantity, total)
		if err != nil {
			log.Println("Failed to create transaction:", err)
			return
		}

		log.Println("---------------------------")
		log.Printf("Transaction created successfully with details:\nInvoice\t\t: %s\nTransdate\t: %s\nProduct ID\t: %d\nCustomer ID\t: %d\nQuantity\t: %d\nTotal\t\t: %d\nCreated by\t: %d\n", transaction.Invoice, transaction.Transdate, transaction.ProductID, transaction.CustomersID, transaction.Quantity, transaction.Total, transaction.CreatedBy)
		log.Println("---------------------------")
	}
}
