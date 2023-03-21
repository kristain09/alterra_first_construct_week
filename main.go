package main

import (
	"first_construct_week/config"
	"first_construct_week/products"
	"fmt"
	"log"
)

func main() {
	
	running := true
	var choice int
	var username string
	var password string
	var choice2 int
	//db
	// fmt.Scanln()
	for running {
		fmt.Println("Welcome to our project!\nwhat do you want to do?")
		fmt.Println("1. Login")
		fmt.Println("99. Exit")
		fmt.Print("Input menu : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Please enter your username!")
			fmt.Scan(&username)
			fmt.Println("Please enter your password!")
			fmt.Scan(&password)
			// function login
			//
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
				login = false
			case 99:
				running = false
			default:
				fmt.Println("anda memasukkan input yang salah!")
				continue
			}
		}
	}
}

//product :
//product has transaction jadi 1

// user: admin
//login dan autentifikasi admin!
// user sama customer

// transaction bagi 2 di akhir

func inputtransaksi() {

}
