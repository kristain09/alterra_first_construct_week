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

func (cm CustomerModels) GetAllCustomerData(id int) ([]Customer, error) {
	rows, err := cm.conn.Query("SELECT * FROM customers")
	if err != nil {
		log.Println("Error executing query:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.UserID); err != nil {
			log.Println("Error scanning row:", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating rows:", err.Error())
		return nil, err
	}

	return customers, nil
}
