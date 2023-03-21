package products

import (
	"database/sql"
	"fmt"
)

type ProductModel struct {
	conn *sql.DB
}

func (um *ProductModel) SetConnection(db *sql.DB) {
	um.conn = db
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
