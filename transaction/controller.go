package transaction

import (
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
	TrCustController   *customer.CustomerController
	TrProdModels       products.ProductModel
	TrUsersModels      users.UsersModels
}

func (tc *TransactionsController) SetConnTcTrModels(tm TransactionsModels) {
	tc.TransactionsModels = tm
}

func (tc *TransactionsController) SetConnTcCustomer(cc customer.CustomerController) {
	tc.TrCustController = &cc
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

func (tc *TransactionsController) CreateTransaction(id int) error { // id login
	var choice int
	fmt.Println("1. New Customer")
	fmt.Println("2. Existing Customer")
	fmt.Scanln(&choice)

	var custConn TransactionsController
	switch choice {
	case 1:
		//call RegisterCustomer customer package!

		cust_id, err := custConn.TrCustController.RegisterCustomer(id)
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

		// print all customer by id
		custConn.TrCustController.PrintAllCustomerData(id)
		fmt.Println("Enter customer's ID")
		fmt.Scanln(&trInput.CustomersID)

		var productList []products.Products
		choie := 1
		for choice != 0 {
			var list products.ProductController
			list.HandleListProduct()
			fmt.Println("Enter product's id\nPress 0 To Exit")
			fmt.Scanln(&productName.ID)
			trInput.Product = append(trInput.Product, productName)

			fmt.Println("Enter qty")
			fmt.Scanln(&quantity)
		}
	default:
		fmt.Println("Not valid operation")
		return nil
	}
}
