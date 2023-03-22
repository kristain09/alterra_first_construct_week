package customer

import "database/sql"

type CustomerModels struct {
	conn *sql.DB
}

func (cm *CustomerModels) SetConnCustModels(db *sql.DB) {
	cm.conn = db
}
