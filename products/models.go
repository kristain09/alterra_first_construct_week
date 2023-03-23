package products

import (
	"database/sql"
	"fmt"
)

type ProductModel struct {
	conn *sql.DB
}

func (pm *ProductModel) SetConnection(db *sql.DB) {
	pm.conn = db
}

func (pm *ProductModel) ListProduct(name string, price int, stock int, updatedAt string, createdBy int) ([]Products, error) {
	var products []Products

	query := "SELECT products.id, products.name, products.price, products.stock, products.created_by, products.deleted_at, products.updated_at, users.id as user_id, users.username FROM products JOIN users ON products.created_by = users.id WHERE products.deleted_at IS NULL"

	// args := []interface{}{}

	// bisa menerapkan filter var arguments menggunakan interface kosong dengan input dari users.

	rows, err := pm.conn.Query(query)
	if err != nil {
		return products, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Products
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Created_by, &p.Deleted_at, &p.Updated_at, &p.Created_by, &p.Username)
		if err != nil {
			return products, fmt.Errorf("error scanning row: %w", err)
		}
		products = append(products, p)
	}

	err = rows.Err()
	if err != nil {
		return products, fmt.Errorf("error iterating over rows: %w", err)
	}

	return products, nil
}

func (pm *ProductModel) CreateProduct(name string, price int, stock int, createdBy int) (Products, error) {
	query := "INSERT INTO products (name, price, stock, created_by) VALUES (?, ?, ?, ?)"
	result, err := pm.conn.Exec(query, name, price, stock, createdBy)
	if err != nil {
		return Products{}, fmt.Errorf("error creating product: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Products{}, fmt.Errorf("error getting product ID: %w", err)
	}

	// Retrieve the details of the newly created product from the database using the ID of the newly inserted row
	query = "SELECT id, name, price, stock, created_by, deleted_at, updated_at FROM products WHERE id = ?"
	row := pm.conn.QueryRow(query, id)
	product := Products{}
	err = row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.Created_by, &product.Deleted_at, &product.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return Products{}, fmt.Errorf("product with ID %d not found", id)
		}
		return Products{}, fmt.Errorf("error getting product details: %w", err)
	}
	return product, nil
}
