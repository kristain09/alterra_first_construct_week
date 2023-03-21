package config

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:pilotjhon18@tcp(localhost:3306)/mydb")
	if err != nil {
		t.Fatalf("Error opening database connection: %s", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Error pinging database: %s", err.Error())
	}
}

