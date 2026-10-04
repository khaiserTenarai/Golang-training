package repository

import (
	"database/sql"
	"fmt"
	"task10_order_inventory_transaction/models"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateProduct(p models.Product) (int, error) {
	var id int
	err := r.DB.QueryRow(
		`INSERT INTO inv_products (name, price, stock) VALUES ($1, $2, $3) RETURNING id`,
		p.Name, p.Price, p.Stock).Scan(&id)
	return id, err
}

func (r *OrderRepository) GetAllProducts() ([]models.Product, error) {
	rows, err := r.DB.Query(`SELECT id, name, price, stock, created_at FROM inv_products ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

// PlaceOrder uses a transaction to create order and reduce stock; rolls back if stock is insufficient
func (r *OrderRepository) PlaceOrder(productID, quantity int) (int, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, fmt.Errorf("begin transaction: %w", err)
	}

	// Lock the product row and check stock
	var stock int
	var price float64
	var productName string
	err = tx.QueryRow(
		`SELECT name, price, stock FROM inv_products WHERE id=$1 FOR UPDATE`, productID).
		Scan(&productName, &price, &stock)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("product not found: %w", err)
	}

	// Check if stock is sufficient
	if stock < quantity {
		tx.Rollback()
		// Log failed order
		r.DB.Exec(
			`INSERT INTO inv_orders (product_id, quantity, total_price, status) VALUES ($1, $2, $3, 'FAILED')`,
			productID, quantity, float64(quantity)*price)
		return 0, fmt.Errorf("insufficient stock for '%s': available %d, requested %d (ROLLED BACK)", productName, stock, quantity)
	}

	// Reduce stock
	_, err = tx.Exec(`UPDATE inv_products SET stock = stock - $1 WHERE id = $2`, quantity, productID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("reduce stock: %w", err)
	}

	totalPrice := float64(quantity) * price

	// Create the order record
	var orderID int
	err = tx.QueryRow(
		`INSERT INTO inv_orders (product_id, quantity, total_price, status) VALUES ($1, $2, $3, 'SUCCESS') RETURNING id`,
		productID, quantity, totalPrice).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("create order: %w", err)
	}

	// COMMIT
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("commit: %w", err)
	}

	fmt.Printf("[TX] Order committed: %s x%d = %.2f | Stock: %d -> %d\n",
		productName, quantity, totalPrice, stock, stock-quantity)
	return orderID, nil
}

func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	rows, err := r.DB.Query(
		`SELECT o.id, o.product_id, p.name, o.quantity, o.total_price, o.status, o.created_at
		FROM inv_orders o
		JOIN inv_products p ON o.product_id = p.id
		ORDER BY o.created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.ProductID, &o.ProductName, &o.Quantity, &o.TotalPrice, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, rows.Err()
}
