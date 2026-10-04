package repository

import (
	"database/sql"
	"fmt"
	"task09_ecommerce_order_system/models"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(customerID int, items []models.OrderItem) (int, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, fmt.Errorf("begin transaction: %w", err)
	}

	var totalAmount float64
	for i := range items {
		items[i].Subtotal = float64(items[i].Quantity) * items[i].UnitPrice
		totalAmount += items[i].Subtotal
	}

	var orderID int
	err = tx.QueryRow(
		`INSERT INTO ecom_orders (customer_id, total_amount) VALUES ($1, $2) RETURNING id`,
		customerID, totalAmount).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("create order: %w", err)
	}

	for _, item := range items {
		_, err = tx.Exec(
			`INSERT INTO ecom_order_items (order_id, product_id, quantity, unit_price, subtotal) VALUES ($1, $2, $3, $4, $5)`,
			orderID, item.ProductID, item.Quantity, item.UnitPrice, item.Subtotal)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("add order item: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("commit: %w", err)
	}
	return orderID, nil
}

// GetOrderDetails uses JOIN queries to display complete order details
func (r *OrderRepository) GetOrderDetails(orderID int) (models.Order, error) {
	var order models.Order

	// Get order with customer name using JOIN
	err := r.DB.QueryRow(
		`SELECT o.id, o.customer_id, c.name, o.total_amount, o.status, o.created_at
		FROM ecom_orders o
		JOIN ecom_customers c ON o.customer_id = c.id
		WHERE o.id = $1`, orderID).
		Scan(&order.ID, &order.CustomerID, &order.CustomerName, &order.TotalAmount, &order.Status, &order.CreatedAt)
	if err != nil {
		return order, err
	}

	// Get order items with product names using JOIN
	rows, err := r.DB.Query(
		`SELECT oi.id, oi.order_id, oi.product_id, p.name, oi.quantity, oi.unit_price, oi.subtotal
		FROM ecom_order_items oi
		JOIN ecom_products p ON oi.product_id = p.id
		WHERE oi.order_id = $1`, orderID)
	if err != nil {
		return order, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.ProductName,
			&item.Quantity, &item.UnitPrice, &item.Subtotal); err != nil {
			return order, err
		}
		order.Items = append(order.Items, item)
	}
	return order, rows.Err()
}

// GetAllOrders uses JOIN to list all orders with customer names
func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	rows, err := r.DB.Query(
		`SELECT o.id, o.customer_id, c.name, o.total_amount, o.status, o.created_at
		FROM ecom_orders o
		JOIN ecom_customers c ON o.customer_id = c.id
		ORDER BY o.created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.CustomerID, &o.CustomerName, &o.TotalAmount, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, rows.Err()
}
