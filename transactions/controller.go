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
			fmt.Println(" Product Information\n", now.Format("Monday, 2006 January 2 15:04:05"))
			pm := products.ProductModel{}
			pm.SetConnection(connection)
			pc := products.NewProductController(&pm)
			if connection == nil {
				log.Fatalln(" connected")
			}
			pc.HandleListProduct()
			tc.CreateTransaction
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

func (tc *TransactionsController) CreateTransaction(id int) error { // id login
	var choice int
	fmt.Println("1. New Customer")
	fmt.Println("2. Existing Customer")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		//call RegisterCustomer customer package!
		var custConn TransactionsController
		cust_id, err := custConn.TrCustModels.RegisterCustomer(id)
		if err != nil {
			log.Print(err)
			return nil //gatau bener apa salah
		}
		return nil //kayanya bukan return tpi nanti dipake dibawah
	case 2:
		var trInput Transactions
		var (
			productName products.Products
			quantity    int
		)
		invoice, err := strconv.Atoi(time.Now().Format("020106"))
		if err != nil {
			log.Print("Fail to generate invoice", err.Error())
		}
		trInput.Invoice = invoice
		// input transaction
		// print all customer by id

		fmt.Println("Enter customer's ID")
		fmt.Scanln(&trInput.CustomersID)

		var productList []products.Products
		var choice int
		for choice > 0 {
			fmt.Println("Enter product's id")
			fmt.Scanln(&productName.ID)
			
			product, err := products.
			trInput.Product = append(trInput.Product, productName)
			fmt.Println("Enter qty")
			fmt.Scanln(&quantity)
		}
	default:
		fmt.Println("Not valid operation")
		return nil
	}
}