package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:sasi2356@localhost:5432/assignment_db"

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	// Schema Setup
	_, _ = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS store_customers (id SERIAL PRIMARY KEY, name VARCHAR(100));
		CREATE TABLE IF NOT EXISTS store_products (id SERIAL PRIMARY KEY, name VARCHAR(100), price NUMERIC(10,2));
		CREATE TABLE IF NOT EXISTS orders (id SERIAL PRIMARY KEY, customer_id INT REFERENCES store_customers(id));
		CREATE TABLE IF NOT EXISTS order_items (id SERIAL PRIMARY KEY, order_id INT REFERENCES orders(id), product_id INT REFERENCES store_products(id), quantity INT);
	`)

	// Data insertion
	var custID, prodID, orderID int
	_ = conn.QueryRow(ctx, "INSERT INTO store_customers (name) VALUES ('Ivan') RETURNING id").Scan(&custID)
	_ = conn.QueryRow(ctx, "INSERT INTO store_products (name, price) VALUES ('Wireless Mouse', 25.00) RETURNING id").Scan(&prodID)
	_ = conn.QueryRow(ctx, "INSERT INTO orders (customer_id) VALUES ($1) RETURNING id", custID).Scan(&orderID)
	_, _ = conn.Exec(ctx, "INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)", orderID, prodID, 2)

	// JOIN Query
	query := `
		SELECT o.id, c.name, p.name, oi.quantity, p.price, (oi.quantity * p.price) as total
		FROM orders o
		JOIN store_customers c ON o.customer_id = c.id
		JOIN order_items oi ON o.id = oi.order_id
		JOIN store_products p ON oi.product_id = p.id
		WHERE o.id = $1`

	rows, err := conn.Query(ctx, query, orderID)
	if err != nil {
		log.Fatalf("Order query failed: %v\n", err)
	}
	defer rows.Close()

	fmt.Println("--- Detailed Order Receipt ---")
	for rows.Next() {
		var oID, qty int
		var custName, prodName string
		var price, total float64
		_ = rows.Scan(&oID, &custName, &prodName, &qty, &price, &total)
		fmt.Printf("Order #%d | Customer: %s | Item: %s | Qty: %d | Unit Price: $%.2f \vert{} Total:$%.2f\n", oID, custName, prodName, qty, price, total)
	}
}