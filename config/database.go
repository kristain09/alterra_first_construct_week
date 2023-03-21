package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectSql(a appConfig) *sql.DB {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		a.DBUserName, a.dbPassword, a.DBHost, a.DBPort, a.DBName)

	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		log.Fatal("unable to connect to database! please try again!\n", err.Error())
		return nil
	}
	return db
}
