package config

import (
	"database/sql"
	"fmt"
	"log"
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
		log.Fatal("unable to connect to database! please try again!\n", err.Error())
		return nil
	}
	return db
}
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
