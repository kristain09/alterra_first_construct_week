package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(config DatabaseConfig) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	db.SetMaxIdleConns(5)

	db.SetMaxOpenConns(50)

	db.SetConnMaxIdleTime(5 * time.Minute)

	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to establish a good connection to the database: %v", err)
	}

	return db, nil
}
