package customer

import (
	"database/sql"
	"log"
)

type CustomerModels struct {
	conn *sql.DB
}

func (cm *CustomerModels) SetConnDBCustModels(db *sql.DB) {
	cm.conn = db
}

func (cm *CustomerModels) InsertDataToCustomers(newCustomer Customer, id int) (int, error) { //jgn lupa if kembalian 0

	result, err := cm.conn.Exec("Insert into customers (name, user_id) values (?)", newCustomer.Name, id)

	if err != nil {
		log.Println("Error executing", err.Error())
		return 0, err
	}
	resultAff, err := result.RowsAffected()

	if err != nil {
		log.Println("error getting rows affected", err.Error())
		return 0, err
	}

	if resultAff <= 0 {
		log.Println("now rows are affected", err.Error())
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return 0, err
	}
	newCustomer.ID = int(lastInsertID)

	return newCustomer.ID, nil
}
