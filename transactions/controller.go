package transactions

import (
	customer "first_construct_week/customers"
	"first_construct_week/products"
	"first_construct_week/users"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type TransactionsController struct {
	TransactionsModels TransactionsModels
	TrCustModels       customer.CustomerController
	TrProdModels       products.ProductModel
	TrUsersModels      users.UsersModels
}

func NewTransactionController(tm *TransactionsModels) *TransactionsController {
	return &TransactionsController{
		TransactionsModels: TransactionsModels{},
		TrCustModels:       customer.CustomerController{},
		TrProdModels:       products.ProductModel{},
		TrUsersModels:      users.UsersModels{},
	}
}

func (tc *TransactionsController) SetConnTcTrModels(tm TransactionsModels) {
	tc.TransactionsModels = tm
}

func (tc *TransactionsController) SetConnTcCustomer(cc customer.CustomerController) {
	tc.TrCustModels = cc
}

func (tc *TransactionsController) HandleRequest() {
	var choice int
	for {
		fmt.Println("1. Transaction History\n2. Delete Transaction\n3. Back")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			tc.TransactionHistory()
		case 2:
			tc.DeleteTransaction()
		case 3:
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
