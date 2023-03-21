package main

import (
	"first_construct_week/config"
	"first_construct_week/users"
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
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			user, err := ctr.Login()
			if err != nil {
				log.Print(err)
				continue
			
			fmt.Println(user)
			//
			//
			//
			//

		case 99:
			running = false
		default:
			fmt.Println("incorrect input!\nPlease try again!")
			continue
		}
		logIn := true
		for logIn {
			menu1 := `
1. Product Information
2. Transaction Input
3. Transaction History
4. Register Cashier
9. Logout
99. Exit`
			fmt.Println(menu1)
			fmt.Scan(&choice2)
			switch choice2 {
			case 1:
				// method atau function barang
				// query hanya di model query exec
				// function/method ada di entities
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
				// method atau function register kasir
				// if branching first!
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
}
