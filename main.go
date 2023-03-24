package main

import (
	"first_construct_week/config"
	customer "first_construct_week/customers"
	"first_construct_week/products"
	"first_construct_week/transaction"
	"first_construct_week/users"
	"fmt"
	"log"
)

func main() {
	running := true
	var (
		choice  int
		choice2 int
		um      users.UsersModels
		uc      users.UsersController
		auth    users.Users
		pm      products.ProductModel
		tc      transaction.TransactionsController
		tm      transaction.TransactionsModels
		cc      customer.CustomerController
		cm      customer.CustomerModels
	)

	cfg := config.InitConfig()
	connection := config.ConnectSql(*cfg)
	defer connection.Close()

	if connection == nil {
		log.Fatal("error connection to database")
	}

	//membuat koneksi antar controller dan models
	um.SetConnUsersModels(connection)
	pm.SetConnection(connection)
	uc.SetConnectModels(um)
	pc := products.NewProductController(&pm)
	tm.SetConnDBTransModels(connection)
	tc.SetConnTcTrModels(tm)
	cm.SetConnDBCustModels(connection)
	cc.SetConnCcCustModels(cm)
	tc.TrCustController = cc
	tc.TrProdModels = pm
	tc.TrUsersModels = um

	for running {
		fmt.Println("Welcome to our project!")
		fmt.Println("what do you want to do")
		fmt.Println("1. Login")
		fmt.Println("99. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			result, err := uc.Login()
			auth = *result
			if err != nil {
				log.Print(err)
				continue
			}
			fmt.Printf("Halo %s, What can we do for you?\n", auth.UserName)
		case 99:
			running = false
		default:
			fmt.Println("incorrect input!\nPlease try again!")
			continue
		}

		LoggedIn := true

		if auth.UserName == "admin" {
			for LoggedIn {
				menu := "1. Product Information\n2. Transaction Input\n3. Delete Transaction\n5. All Transaction History\n6.AllTransaction By Cashier\n7. Register Cashier\n8. Delete Cashier\n9. Logout\n99. Exit\n"
				fmt.Println(menu)
				fmt.Scanln(&choice2)
				switch choice2 {
				case 1:
					pc.HandleRequest(auth.ID) //done
				case 2:
					tc.CreateTransaction(auth.ID) //done
				case 3:
					tc.DeleteTransaction()
				case 5:
					tc.TransactionHistory()
				case 6:
					tc.TransactionHistoryByID(auth.ID)
				case 7:
					uc.Register() //done
				case 8:
					uc.DeleteUser() //done
				case 9:
					LoggedIn = false
				case 99:
					LoggedIn = false
					running = false
				default:
					fmt.Println("Incorrect input! Please try again!")
				}
			}
		} else {
			for LoggedIn {
				menu := "1. Product Information\n2. Transaction Input\n3. Transaction History\n9. Logout\n99. Exit\n"
				fmt.Println(menu)
				fmt.Scanln(&choice2)
				switch choice2 {
				case 1:
					pc.HandleRequest(auth.ID)
				case 2:
					tc.CreateTransaction(auth.ID)
				case 3:
					tc.TransactionHistoryByID(auth.ID)
				case 9:
					LoggedIn = false
				case 99:
					LoggedIn = false
					running = false
				default:
					fmt.Println("Incorrect input! Please try again!")
				}
			}
		}
	}
}
