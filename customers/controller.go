package customer

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
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

func (cc CustomerController) PrintAllCustomerData(id int) error {

	customer, err := cc.CustomerModels.GetAllCustomerData(id)
	if err != nil {
		log.Print(err)
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Created_by"})

	for _, cust := range customer {
		row := []string{
			strconv.Itoa(cust.ID),
			cust.Name,
			strconv.Itoa(cust.UserID),
			"-",
		}
		table.Append(row)
	}

	table.Render()
	return nil
}
