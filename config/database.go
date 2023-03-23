package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(databaseConfig DatabaseConfig) *sql.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Name)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("unable to connect to database! please try again!\n", err.Error())
		return nil
	}
	return db
}
