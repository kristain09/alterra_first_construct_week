package transactions

import (
	"first_construct_week/config"
	"first_construct_week/products"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type TransactionController struct {
	transactionModel *TransactionsModels
}

func NewTransactionController(tm *TransactionsModels) *TransactionController {
	return &TransactionController{tm}
}

func (tc *TransactionController) HandleRequest() {
	var choice int
	for {
		var id int
		fmt.Println("1. Transaction History\n2. Create Transaction\n3. Delete Transaction\n4.Back")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Print("Enter customer ID: ")
			fmt.Scanln(&id)
			tc.TransactionHistoryByID(id)
		case 2:
			cfg := config.InitConfig()
			connection, _ := config.GetConnection(*cfg)
			now := time.Now()
			fmt.Println(" Transaction Information\n", now.Format("Monday, 2006 January 2 15:04:05"))
			pm := products.ProductModel{}
			pm.SetConnection(connection)
			pc := products.NewProductController(&pm)
			if connection == nil {
				log.Fatalln(" connected")
			}
			pc.HandleListProduct()
			tc.handleCreateTransaction()
		case 3:
			tc.DeleteTransaction()
		case 4:
			return
		default:
			fmt.Println("Incorrect input! Please try again!")
		}
	}
}

func (tc TransactionController) TransactionHistoryByID(id int) error {
	transactions, err := tc.transactionModel.GetAllTransDataByID(id)
	if err != nil {
		log.Print(err)
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Invoice", "Transdate", "Quantity", "Total", "Customers ID", "Created By"})

	for _, transaction := range transactions {
		row := []string{
			fmt.Sprintf("%v", transaction.Invoice),
			fmt.Sprintf("%v", transaction.TransDate),
			fmt.Sprintf("%d", transaction.Quantity),
			fmt.Sprintf("%d", transaction.Total),
			strconv.Itoa(transaction.CustomersID),
			strconv.Itoa(transaction.CreatedBy),
		}
		table.Append(row)
	}

	table.Render()
	return nil
}

func (tc TransactionController) DeleteTransaction() {
	var id int
	fmt.Println("Please enter transaction id!")
	fmt.Scanln(&id)
	tc.transactionModel.InitDeletedAt(id)
}

func (tc *TransactionController) handleCreateTransaction() {
	var customersID, createdBy, quantity, total int

	fmt.Println("---------------------------")
	fmt.Println("Create Transaction")
	fmt.Println("---------------------------")

	for {
		fmt.Print("Input products? (y/n): ")
		var choice string
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}
		if choice != "y" {
			break
		}

		var productID int
		fmt.Print("Enter product ID: ")
		if _, err := fmt.Scan(&productID); err != nil {
			log.Println("Invalid input, please try again.")
			continue
		}

		fmt.Print("Enter quantity: ")
		if _, err := fmt.Scan(&quantity); err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		fmt.Print("Enter total: ")
		if _, err := fmt.Scan(&total); err != nil {
			log.Println("Invalid input, please try again.")
			continue
		}

		fmt.Print("Enter customer ID: ")
		fmt.Scan(&customersID)

		fmt.Print("Enter created by user ID: ")
		fmt.Scan(&createdBy)

		// total := product.Price * quantity

		transaction, err := tc.transactionModel.CreateTransaction(customersID, createdBy, quantity, total)
		if err != nil {
			log.Println("Failed to create transaction:", err)
			return
		}

		fmt.Println("---------------------------")
		fmt.Printf("Transaction created successfully with details:\nInvoice\t\t: %s\nTransdate\t: %s\nProduct ID\t: %d\nCustomer ID\t: %d\nQuantity\t: %d\nTotal\t\t: %d\nCreated by\t: %d\n", transaction.Invoice, transaction.TransDate, transaction.Product, transaction.CustomersID, transaction.Quantity, transaction.Total, transaction.CreatedBy)
		log.Println("---------------------------")
	}
}
