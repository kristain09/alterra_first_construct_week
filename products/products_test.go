package products

import (
	"database/sql"
	"first_construct_week/config"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestExecSql(t *testing.T) {
	conn := config.InitDatabase()
	db, _ := config.GetConnection(*conn)
	defer db.Close()

	query := "INSERT INTO products (name, price, stock, deleted_at, created_by) VALUES ('Cap Panda', '100000', '100','2022-12-31 23:59:59', '1');"
	res, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error executing SQL query: %s", err.Error())
	}

	log.Println("Success insert new product")

	aff, err := res.RowsAffected()
	assert.Greater(t, aff, int64(0))
	assert.Nil(t, err)
}

func TestQuerySql(t *testing.T) {
	conn := config.InitDatabase()
	db, _ := config.GetConnection(*conn)
	defer db.Close()

	query := "SELECT id, name, price, stock, deleted_at, created_by FROM products"
	rows, err := db.Query(query)
	if err != nil {
		t.Fatalf("Error querying database: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id, price, stock, created_by int
		var name string
		var deleted_at sql.NullString
		err := rows.Scan(&id, &name, &price, &stock, &deleted_at, &created_by)
		if err != nil {
			log.Fatalf("Error scanning row: %s", err.Error())
		}
		log.Println("=====================================")
		log.Println("Id:", id)
		log.Println("Name:", name)
		log.Println("Price:", price)
		log.Println("Stock:", stock)
		if deleted_at.Valid {
			log.Println("Deleted_at:", deleted_at)
		}
		log.Println("Created_at", created_by)
	}
}

// 	for rows.Next() {
// 		var id, price, stock, created_by int
// 		var name string
// 		var deleted_at sql.NullTime
// 		err := rows.Scan(&id, &name, &price, &stock, &deleted_at, &created_by)
// 		if err != nil {
// 			t.Fatalf("Error scanning row: %s", err.Error())
// 		}
// 		fmt.Println("=====================================")
// 		fmt.Println("Id:", id)
// 		fmt.Println("Name:", name)
// 		fmt.Println("Price:", price)
// 		fmt.Println("Stock:", stock)
// 		if deleted_at.Valid {
// 			fmt.Println("Deleted_at:", deleted_at.Time)
// 		}
// 		fmt.Println("Created_at", created_by)
// 	}
// }
