package transaction

import (
	"errors"
	customer "first_construct_week/customers"
	"first_construct_week/products"
	"first_construct_week/users"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type TransactionsController struct {
	TransactionsModels TransactionsModels
	TrCustController   customer.CustomerController
	TrProdModels       products.ProductModel
	TrUsersModels      users.UsersModels
}

func (tc *TransactionsController) SetConnTcTrModels(tm TransactionsModels) {
	tc.TransactionsModels = tm
}

func (tc *TransactionsController) SetConnTcCustomer(cc customer.CustomerController) {
	tc.TrCustController = cc
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
			transaction.TransDate,
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
			transaction.TransDate,
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

func (tc TransactionsController) DeleteTransaction() error {
	var id int
	var err error
	fmt.Println("Please enter transaction id!")
	_, err = fmt.Scanln(&id)
	if err != nil {
		return err
	}
	err = tc.TransactionsModels.InitDeletedAt(id)
	if err != nil {
		return err
	}
	return nil
}

func (tc *TransactionsController) CreateTransaction(id int) error { // id login
	var choice int
	fmt.Println("1. New Customer")
	fmt.Println("2. Existing Customer")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		// call RegisterCustomer customer package!

		cust_id, err := tc.TrCustController.RegisterCustomer(id)
		if err != nil {
			log.Print(err)
			return err
		}
		var trInput Transactions
		var (
			productName products.Products
			quantity    int
		)
		timeSeed := time.Now().UnixNano()
		rand.Seed(timeSeed)
		invoice := rand.Intn(10000000)
		if err != nil {
			log.Print("Fail to generate invoice", err.Error())
		}
		trInput.Invoice = invoice
		trInput.CustomersID = cust_id

		var productList []products.Products
		choice := 1
		for choice != 0 {
			productList, err = tc.TrProdModels.ListProduct("", 0, 0, 0)
			if err != nil {
				return err
			}
			for _, p := range productList {
				fmt.Println(fmt.Sprint(p.ID), p.Name, fmt.Sprintf("%d", p.Price), fmt.Sprintf("%d", p.Stock), p.Username)
			}
			fmt.Println("Enter product's id\nPress 0 To Exit")
			fmt.Scanln(&productName.ID)
			if productName.ID == 0 {
				break
			}

			fmt.Println("Enter qty")
			fmt.Scanln(&quantity)
			product, err := tc.TrProdModels.GetProductByID(productName.ID)
			if err != nil {
				log.Print(err)
				continue
			}

			if product.Stock < quantity {
				fmt.Println("Not Enough Stock")
				continue
			}
			productList = append(productList, products.Products{
				ID:         product.ID,
				Name:       product.Name,
				Price:      product.Price,
				Stock:      quantity,
				Created_by: product.Created_by,
				Username:   product.Username,
			})
		}
		trInput.Product = productList

		_, err = tc.TransactionsModels.InsertTransaction(trInput, id)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Transaction Successfully Created")
		return nil

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
		tc.TrCustController.PrintAllCustomerData(id)
		fmt.Println("Enter customer's ID")
		fmt.Scanln(&trInput.CustomersID)

		var productList []products.Products
		choice := 1
		for choice != 0 {
			productList, err = tc.TrProdModels.ListProduct("", 0, 0, 0)
			if err != nil {
				return err
			}
			for _, p := range productList {
				fmt.Println(fmt.Sprint(p.ID), p.Name, fmt.Sprintf("%d", p.Price), fmt.Sprintf("%d", p.Stock), p.Username)
			}
			fmt.Println("Enter product's id\nPress 0 To Exit")
			fmt.Scanln(&productName.ID)
			if productName.ID == 0 {
				break
			}

			fmt.Println("Enter qty")
			fmt.Scanln(&quantity)
			product, err := tc.TrProdModels.GetProductByID(productName.ID)
			if err != nil {
				log.Print(err)
				continue
			}

			if product.Stock < quantity {
				fmt.Println("Not Enough Stock")
				continue
			}
			productList = append(productList, products.Products{
				ID:         product.ID,
				Name:       product.Name,
				Price:      product.Price,
				Stock:      quantity,
				Created_by: product.Created_by,
				Username:   product.Username,
			})
		}
		trInput.Product = productList

		_, err = tc.TransactionsModels.InsertTransaction(trInput, id)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Transaction Successfully Created")
		return nil
	default:
		return errors.New("invalid choice, please enter valid choice")
	}
}