//file untuk konfigurasi database secara umum bagi aplikasi.

package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct{
	Host			string
	Port			string
	Username 	string
	Password 	string
	DBName 		string
}

func InitDatabase() (*sql.DB, error) {
	databaseConfig := DatabaseConfig {
		Host: "localhost",
		Port: "3306",
		Username: "root",
		Password: "@Pilotjhon18",
		DBName: "mydb",
	}
	
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
		databaseConfig.Username, 
		databaseConfig.Password, 
		databaseConfig.Host, 
		databaseConfig.Port, 
		databaseConfig.DBName))

	if err != nil {
		log.Fatal("error connecting to database: ", err)
		return nil, err
	}
	//test database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// fmt.Println("Database connection established")

	return db, nil
}