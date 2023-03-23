package main

import (
	"first_construct_week/config"
	"first_construct_week/products"
	transaction "first_construct_week/transactions"
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
	)

	cfg := config.InitConfig()
	connection, _ := config.GetConnection(*cfg)
	defer connection.Close()

	if connection == nil {
		log.Fatal("error connection to database")
	}

	um.SetConnUsersModels(connection)
	pm.SetConnection(connection)
	uc.SetConnectModels(um)
	pc := products.NewProductController(&pm)

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
				menu := "1. Product Information\n2. Transaction Input\n3. Update Transaction\n4. Delete Transaction\n5. All Transaction History\n6. Register Cashier\n7. Delete Cashier\n9. Logout\n99. Exit\n"
				fmt.Println(menu)
				fmt.Scanln(&choice2)
				switch choice2 {
				case 1:
					// Product Information functionality
					pc.HandleRequest()
				case 2:
					//  Transaction Input functionality
				case 3:
					// Update Transaction History functionality
				case 4:
					// template m c//Delete Transaction functionality
				case 5:
					// template m c// All Transaction History functionality
				case 6:
					// template m c// Transaction By ID
				case 7:
					//done // Register Cashier functionality
					uc.Register()
				case 8:
					//done // Delete Cashier functionality
					uc.DeleteUser()
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
					// Product Information functionality
					pm := products.ProductModel{}
					pm.SetConnection(connection)
					pc := products.NewProductController(&pm)
					if connection == nil {
						log.Fatalln(" connected")
					}
					pc.HandleRequest()
				case 2:
					// Transaction Input functionality
					tm := transaction.TransactionsModels{}
					tm.SetConnDBTransModels(connection)
					tc := transaction.NewTransactionController(&tm)
					if connection == nil {
						log.Fatalln(" connected")
					}
					tc.HandleRequest()
				case 3:
					// template m c// Transaction History functionality
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
