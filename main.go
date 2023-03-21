package main

import (
	"first_construct_week/config"
	"first_construct_week/users"
	"first_construct_week/products"
	"fmt"
	"log"
)

func main() {
	
	running := true
	var (
		choice  int
		choice2 int
		mdl     users.UsersModels
		ctr     users.UsersController
		auth    users.Users
	)

	cfg := config.InitConfig()
	connection := config.ConnectSql(*cfg)
	defer connection.Close()

	if connection == nil {
		log.Fatal("error to connect to database")
	}

	mdl.SetConnUsersModels(connection)
	ctr.SetConnectModels(mdl)
	for running {
		fmt.Println("Welcome to our project!\nwhat do you want to do?")
		fmt.Println("1. Login")
		fmt.Println("99. Exit")
		fmt.Print("Input menu : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			result, err := ctr.Login()
			auth = *result
			if err != nil {
				log.Print(err)
				continue
			}

			//
			//
			//

		case 99:
			running = false
		default:
			fmt.Println("Incorrect input, Please try again!")
			continue
		}
		login := true
		for login {
			fmt.Println("========================")
			fmt.Println("Hi <username>, input menu :")
			fmt.Println("========================")
			conn := config.InitDatabase()
			db, err := config.GetConnection(*conn)
			if err != nil {
				log.Panic(err)
			}
				defer db.Close()

			fmt.Println("1. Product Information")
			fmt.Println("2. Transaction Input")
			fmt.Println("3. Transaction History")
			fmt.Println("4. Register Cashier")
			fmt.Println("9. Logout")
			fmt.Print("Input menu : ")
			fmt.Scan(&choice2)
			switch choice2 {
			case 1:
				cfg := config.InitDatabase()
				conn, _ := config.GetConnection(*cfg)
				pm := products.ProductModel{}
				pm.SetConnection(conn)
				pc := products.NewProductController(&pm)
				if conn == nil {
					log.Fatalln(" connected")
				}
				pc.HandleRequest()

			case 2:
				// method atau function transaksi
				//
				//
				//
				//
				//
				//
				//
				//
				//
				//
				//
				//

			case 3:
				// method atau function rekap penjualan
				//
				//
				//
				//
				//
				//
				//
				//
				//
				//
				//
				//
			case 4:
				if auth.UserName == "admin" {
	
				} else {
					fmt.Println("Acces denied! Please call admin!")
				//
				//
				//
				//
				//
				//
				//
				//
				//
				//
				//

		case 9:
			logIn = false
		case 99:
			running = false
		default:
			fmt.Println("anda memasukkan input yang salah!")
			continue
		}
	}
}
