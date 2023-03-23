package customer

import (
	"fmt"
	"log"
)

type CustomerController struct {
	CustomerModels CustomerModels
}

func (cc *CustomerController) SetConnCcCustModels(cm CustomerModels) {
	cc.CustomerModels = cm
}

func (cc *CustomerController) RegisterCustomer(id int) (int, error) { //id dari login
	var NewCustomer Customer
	fmt.Println("Please input Customer's Name")
	fmt.Scanln(&NewCustomer.Name)

	cust_id, err := cc.CustomerModels.InsertDataToCustomers(NewCustomer, id)

	if err != nil {
		log.Println(err) // kalau fatal disini gimana ya
		return 0, err
	}
	if cust_id == 0 {
		log.Println("error occured there is nothing to change")
	}

	return cust_id, nil
}
