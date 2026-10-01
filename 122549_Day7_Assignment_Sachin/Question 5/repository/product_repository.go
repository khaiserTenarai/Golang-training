// Package repository has all the direct database access for products.
package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"question5-product-inventory/model"
)

// ErrProductNotFound is returned when a product ID doesn't exist.
var ErrProductNotFound = errors.New("product not found")

// ErrInsufficientStock is returned when trying to decrease stock by
// more than what is currently available.
var ErrInsufficientStock = errors.New("not enough stock to decrease by that amount")

// CreateProduct inserts a new product and returns its generated ID.
func CreateProduct(db *sql.DB, p model.Product) (int, error) {
	var id int
	query := `
		INSERT INTO product (name, sku, category, price, quantity)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	err := db.QueryRow(query, p.Name, p.SKU, p.Category, p.Price, p.Quantity).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create product failed: %w", err)
	}
	return id, nil
}

// GetProduct fetches one product by ID.
func GetProduct(db *sql.DB, id int) (model.Product, error) {
	var p model.Product
	query := `SELECT id, name, sku, category, price, quantity, created_at FROM product WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.SKU, &p.Category, &p.Price, &p.Quantity, &p.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Product{}, fmt.Errorf("get product failed: %w", ErrProductNotFound)
		}
		return model.Product{}, fmt.Errorf("get product failed: %w", err)
	}
	return p, nil
}

// ListProducts returns every product, ordered by name.
func ListProducts(db *sql.DB) ([]model.Product, error) {
	query := `SELECT id, name, sku, category, price, quantity, created_at FROM product ORDER BY name`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("list products failed: %w", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.SKU, &p.Category, &p.Price, &p.Quantity, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("list products failed: %w", err)
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

// UpdateProduct updates a product's basic details (not its stock -
// use IncreaseStock/DecreaseStock for that, so every stock change goes
// through the same, safe path).
func UpdateProduct(db *sql.DB, id int, name, category string, price float64) error {
	query := `UPDATE product SET name = $1, category = $2, price = $3 WHERE id = $4`
	result, err := db.Exec(query, name, category, price, id)
	if err != nil {
		return fmt.Errorf("update product failed: %w", err)
	}
	return checkRowsAffected(result, ErrProductNotFound)
}

// DeleteProduct removes a product by ID.
func DeleteProduct(db *sql.DB, id int) error {
	result, err := db.Exec(`DELETE FROM product WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete product failed: %w", err)
	}
	return checkRowsAffected(result, ErrProductNotFound)
}

// IncreaseStock adds `amount` units to a product's quantity. This is a
// single atomic UPDATE statement, so it's already safe even if two
// increases happen for the same product at the same time.
func IncreaseStock(db *sql.DB, id int, amount int) error {
	query := `UPDATE product SET quantity = quantity + $1 WHERE id = $2`
	result, err := db.Exec(query, amount, id)
	if err != nil {
		return fmt.Errorf("increase stock failed: %w", err)
	}
	return checkRowsAffected(result, ErrProductNotFound)
}

// DecreaseStock removes `amount` units from a product's quantity,
// inside a transaction: it locks the row, checks there's enough stock
// left, and only then applies the update. This stops two decreases
// happening at once from pushing quantity below zero.
func DecreaseStock(db *sql.DB, id int, amount int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not start transaction: %w", err)
	}
	defer tx.Rollback()

	var currentQuantity int
	err = tx.QueryRow(`SELECT quantity FROM product WHERE id = $1 FOR UPDATE`, id).Scan(&currentQuantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("decrease stock failed: %w", ErrProductNotFound)
		}
		return fmt.Errorf("decrease stock failed: %w", err)
	}

	if currentQuantity < amount {
		return fmt.Errorf("decrease stock failed: %w", ErrInsufficientStock)
	}

	if _, err := tx.Exec(`UPDATE product SET quantity = quantity - $1 WHERE id = $2`, amount, id); err != nil {
		return fmt.Errorf("decrease stock failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}
	return nil
}

// SearchLowStock returns every product whose quantity is at or below
// the given threshold, lowest quantity first.
func SearchLowStock(db *sql.DB, threshold int) ([]model.Product, error) {
	query := `
		SELECT id, name, sku, category, price, quantity, created_at
		FROM product
		WHERE quantity <= $1
		ORDER BY quantity ASC`
	rows, err := db.Query(query, threshold)
	if err != nil {
		return nil, fmt.Errorf("search low stock failed: %w", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.SKU, &p.Category, &p.Price, &p.Quantity, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("search low stock failed: %w", err)
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func checkRowsAffected(result sql.Result, notFoundErr error) error {
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check affected rows: %w", err)
	}
	if rows == 0 {
		return notFoundErr
	}
	return nil
}
