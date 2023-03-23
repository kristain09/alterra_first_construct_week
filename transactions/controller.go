package transactions

import (
	"first_construct_week/config"
	customer "first_construct_week/customers"
	"first_construct_week/products"
	"first_construct_week/users"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type TransactionsController struct {
	TransactionsModels TransactionsModels
	TrCustModels       customer.CustomerController
	TrProdModels       *products.ProductModel
	TrUsersModels      users.UsersModels
}

func NewTransactionController(tm *TransactionsModels) *TransactionsController {
	return &TransactionsController{
		TransactionsModels: TransactionsModels{},
		TrCustModels:       customer.CustomerController{},
		TrProdModels:       &products.ProductModel{},
		TrUsersModels:      users.UsersModels{},
	}
}

func (tc *TransactionsController) SetConnTcTrModels(tm TransactionsModels) {
	tc.TransactionsModels = tm
}

func (tc *TransactionsController) SetConnTcCustomer(cc customer.CustomerController) {
	tc.TrCustModels = cc
}

func (tc *TransactionsController) SetConnection(pm *products.ProductModel) {
	tc.TrProdModels = pm
}

func (tc *TransactionsController) HandleRequest() {
	var choice int
	for {
		fmt.Println("1. Transaction History\n2. Delete Transaction\n3. Create Transaction\n4.Back")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			tc.TransactionHistory()
		case 2:
			tc.DeleteTransaction()
		case 3:
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
			tc.CreateTransaction()
		case 4:
			return
		default:
			fmt.Println("Incorrect input! Please try again!")
		}
	}
}

func (tc TransactionsController) TransactionHistory() {
	transactions, err := tc.TransactionsModels.GetAllTransData()
	if err != nil {
		log.Print(err)
		return
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
}

func (tc TransactionsController) TransactionHistoryByID(id int) error {
	transactions, err := tc.TransactionsModels.GetAllTransDataByID(id)
	if err != nil {
		log.Print(err)
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

func (tc TransactionsController) DeleteTransaction() {
	var id int
	fmt.Println("Please enter transaction id!")
	fmt.Scanln(&id)
	tc.TransactionsModels.InitDeletedAt(id)
}

func (tc *TransactionsController) CreateTransaction() {
	var productID, customersID, createdBy, quantity, total int

	log.Println("---------------------------")
	log.Println("Create Transaction")
	log.Println("---------------------------")

	for {
		fmt.Print("Add another product? (y/n): ")
		var choice string
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}
		if choice == "n" {
			break
		}

		fmt.Print("Enter product ID: ")
		if _, err := fmt.Scan(&productID); err != nil {
			log.Println("Invalid input, please try again.")
			continue
		}

		fmt.Print("Enter quantity: ")
		if _, err := fmt.Scan(&quantity); err != nil {
			fmt.Println("Invalid input, please try again.")
			continue

			fmt.Print("Enter total: ")
			if _, err := fmt.Scan(&total); err != nil {
				log.Println("Invalid input, please try again.")
				continue
			}
		}

		fmt.Print("Enter customer ID: ")
		fmt.Scan(&customersID)

		fmt.Print("Enter created by user ID: ")
		fmt.Scan(&createdBy)

		// // Query database disini
		// product, err := tc.transactionModel.GetProduct(productID)
		// if err != nil {
		// 	log.Println("Failed to get product:", err)
		// 	return
		// }

		// // Calculate the total by multiplying the price by the quantity
		// total := product.Price * quantity

		transaction, err := tc.TransactionsModels.CreateTransaction(productID, customersID, createdBy, quantity, total)
		if err != nil {
			log.Println("Failed to create transaction:", err)
			return
		}

		fmt.Println("---------------------------")
		fmt.Printf("Transaction created successfully with details:\nInvoice\t\t: %s\nTransdate\t: %s\nProduct ID\t: %d\nCustomer ID\t: %d\nQuantity\t: %d\nTotal\t\t: %d\nCreated by\t: %d\n", transaction.Invoice, transaction.TransDate, transaction.Product, transaction.CustomersID, transaction.Quantity, transaction.Total, transaction.CreatedBy)
		fmt.Println("---------------------------")
	}
}
