package products

import (
	"context"
	"database/sql"
	"first_construct_week/config"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)


func TestExecSql(t *testing.T) {
	db := config.GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	query := "INSERT INTO products (name, price, stock, deleted_at, created_by) VALUES ('Cap Panda', '100000', '100','2022-12-31 23:59:59', '1');"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new product")
}

func TestQuerySql(t *testing.T) {
	db := config.GetConnection()
	defer db.Close()

	ctx := context.Background()

	

	query := "SELECT id, name, price, stock, deleted_at, created_by FROM products"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, price, stock, created_by int
		var name string
		var deleted_at sql.NullTime
		err := rows.Scan(&id, &name, &price, &stock, &deleted_at, &created_by)
		if err != nil {
			panic(err)
		}
		fmt.Println("=====================================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("Price:", price)
		fmt.Println("Stock:", stock)
		if deleted_at.Valid{
			fmt.Println("Deleted_at:", deleted_at.Time)
		}
		fmt.Println("Created_at", created_by)
	}
}

