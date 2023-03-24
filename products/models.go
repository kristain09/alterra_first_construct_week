package products

import (
	"database/sql"
	"fmt"
	"log"
)

type ProductModel struct {
	conn *sql.DB
}

func (pm *ProductModel) SetConnection(db *sql.DB) {
	pm.conn = db
}

func (pm *ProductModel) ListProduct(name string, price int, stock int, updatedAt string, createdBy int) ([]Products, error) {
	var products []Products

	query := "SELECT products.id, products.name, products.price, products.stock, products.created_by, products.deleted_at, DATE_FORMAT(products.updated_at, '%Y-%m-%d %H:%i'), users.id as user_id, users.username FROM products JOIN users ON products.created_by = users.id WHERE products.deleted_at IS NULL ORDER BY products.id, products.updated_at ASC"

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

	// fetch satu baris data product yang baru dibuat user.
	query = "SELECT id, name, price, stock, created_by, updated_at, deleted_at FROM products WHERE id = ?"
	row := pm.conn.QueryRow(query, id)
	product := Products{}
	err = row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.Created_by, &product.Updated_at, &product.Deleted_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return Products{}, fmt.Errorf("product with ID %d not found", id)
		}
		return Products{}, fmt.Errorf("error getting product details: %w", err)
	}
	return product, nil
}

func (pm *ProductModel) UpdateProductNameByID(ID int, createdBy int, name string) (Products, error) {
	query := "UPDATE products SET name = ?, created_by = ?, updated_at = current_timestamp(3) WHERE id = ?"
	result, err := pm.conn.Exec(query, name, createdBy, ID)
	if err != nil {
		log.Println("error updating product name", err.Error())
		return Products{}, err
	}

	resultAff, err := result.RowsAffected()
	if err != nil {
		log.Println("error after updating product name", err.Error())
		return Products{}, err
	}

	if resultAff <= 0 {
		log.Println("error while updating product name", err.Error())
		return Products{}, fmt.Errorf("no rows affected : %w", err)
	}

	//fetch satu baris data product yang baru saja diupdate oleh user.
	query = "SELECT id, name, price, stock, created_by, updated_at, deleted_at FROM products WHERE id = ?"
	row := pm.conn.QueryRow(query, ID)
	product := Products{}
	err = row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.Created_by, &product.Updated_at, &product.Deleted_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return Products{}, fmt.Errorf("product with ID %d not found", ID)
		}
		return Products{}, fmt.Errorf("error getting product details: %w", err)
	}
	return product, nil
}

func (pm *ProductModel) UpdateProductPriceByID(ID int, createdBy int, price int) (Products, error) {
	query := "UPDATE products SET price = ?, created_by = ?, updated_at = current_timestamp(3) WHERE id = ?"
	result, err := pm.conn.Exec(query, price, createdBy, ID)
	if err != nil {
		log.Println("error updating product price", err.Error())
		return Products{}, err
	}

	resultAff, err := result.RowsAffected()
	if err != nil {
		log.Println("error after updating product price", err.Error())
		return Products{}, err
	}

	if resultAff <= 0 {
		log.Println("error while updating product price", err.Error())
		return Products{}, fmt.Errorf("no rows affected : %w", err)
	}

	//fetch satu baris data product yang baru saja diupdate oleh user.
	query = "SELECT id, name, price, stock, created_by, updated_at, deleted_at FROM products WHERE id = ?"
	row := pm.conn.QueryRow(query, ID)
	product := Products{}
	err = row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.Created_by, &product.Updated_at, &product.Deleted_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return Products{}, fmt.Errorf("product with ID %d not found", ID)
		}
		return Products{}, fmt.Errorf("error getting product details: %w", err)
	}
	return product, nil
}

func (pm *ProductModel) UpdateProductStockByID(ID int, createdBy int, stock int) (Products, error) {
	query := "UPDATE products SET stock = ?, created_by = ?, updated_at = current_timestamp(3) WHERE id = ?"
	result, err := pm.conn.Exec(query, stock, createdBy, ID)
	if err != nil {
		log.Println("error updating product stock", err.Error())
		return Products{}, err
	}

	resultAff, err := result.RowsAffected()
	if err != nil {
		log.Println("error after updating product stock", err.Error())
		return Products{}, err
	}

	if resultAff <= 0 {
		log.Println("error while updating product stock", err.Error())
		return Products{}, fmt.Errorf("no rows affected : %w", err)
	}

	//fetch satu baris data product yang baru saja diupdate oleh user.
	query = "SELECT id, name, price, stock, created_by, updated_at, deleted_at FROM products WHERE id = ?"
	row := pm.conn.QueryRow(query, ID)
	product := Products{}
	err = row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.Created_by, &product.Updated_at, &product.Deleted_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return Products{}, fmt.Errorf("product with ID %d not found", ID)
		}
		return Products{}, fmt.Errorf("error getting product details: %w", err)
	}
	return product, nil
}

func (pm *ProductModel) RemoveProductByID(ID int) error {
	query := "UPDATE products SET deleted_at = current_timestamp(3) WHERE id = ?"
	result, err := pm.conn.Exec(query, ID)
	if err != nil {
		log.Println("error removing id product ", err.Error())
		return err
	}

	resultAff, err := result.RowsAffected()
	if err != nil {
		log.Println("error after removing product ", err.Error())
		return err
	}

	if resultAff <= 0 {
		log.Println("error while removing id product ", err.Error())
		return err
	}
	return nil
}

func (pm *ProductModel) GetProductByID(ID int) (*Products, error) {
	var product Products

	query := "SELECT id, name, price, stok, updated_at, created_by, username FROM products WHERE id = ? AND deleted_at IS NULL"
	err := pm.conn.QueryRow(query, ID).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.Updated_at, &product.Created_by, &product.Username)
	if err != nil {
		log.Println("error getting product by id", err.Error())
		return nil, err
	}

	return &product, nil
}
