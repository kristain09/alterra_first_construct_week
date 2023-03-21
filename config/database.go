package config

import (
	"database/sql"
	"fmt"
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
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxIdleConns(50)
	db.SetConnMaxIdleTime(5 *time.Minute)
	db.SetConnMaxLifetime(60 *time.Minute)

	return db
}