package config

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	dataSource := "root:root@tcp(localhost:3306)/mydb?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxIdleConns(50)
	db.SetConnMaxIdleTime(5 *time.Minute)
	db.SetConnMaxLifetime(60 *time.Minute)

	return db
}