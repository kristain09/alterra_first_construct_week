package transaction

import (
	"fmt"
)

type TransactionsController struct {
	TransactionsModels TransactionsModels
}

func (tc *TransactionsController) SetConnTcTrModels(tm TransactionsModels) {
	tc.TransactionsModels = tm
}

func (tc TransactionsController) TransactionHistory() {
	tc.TransactionsModels.PrintAllTransData()
}

func (tc TransactionsController) TransactionHistoryByID(id int) {
	tc.TransactionsModels.PrintTransDataByUserID(id) //khusus admin ada scan di main
}

func (tc TransactionsController) DeleteTransaction() {
	var id int
	fmt.Println("Please enter transaction id!")
	fmt.Scanln(&id)
	tc.TransactionsModels.InitDeletedAt(id)
}

func (tc *TransactionsController) CreateTransaction(id int) { // id login
	var choice int
	fmt.Println("1. New Customer")
	fmt.Println("2. Existing Customer")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		//call RegisterCustomer customer package!
		fallthrough
	case 2:
		// input transaction
		fmt.Println("Please input ")
	default:
		fmt.Println("Not valid operation")
		return
	}
}
