package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const connStr = "postgres://postgres:admin@localhost:5432/assignment_db"

func placeOrder(ctx context.Context, conn *pgx.Conn, customerID, productID, quantity int) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) // Automatic rollback on failure/return

	// 1. Check stock and lock row
	var stock int
	err = tx.QueryRow(ctx, "SELECT stock FROM products WHERE id = $1 FOR UPDATE", productID).Scan(&stock)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	if stock < quantity {
		return fmt.Errorf("insufficient stock: requested %d, available %d", quantity, stock)
	}

	// 2. Reduce stock
	_, err = tx.Exec(ctx, "UPDATE products SET stock = stock - $1 WHERE id = $2", quantity, productID)
	if err != nil {
		return err
	}

	// 3. Create order
	var orderID int
	err = tx.QueryRow(ctx, "INSERT INTO orders (customer_id) VALUES ($1) RETURNING id", customerID).Scan(&orderID)
	if err != nil {
		return err
	}

	// 4. Create order item
	_, err = tx.Exec(ctx, "INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)", orderID, productID, quantity)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}
	defer conn.Close(ctx)

	var prodID, custID int
	_ = conn.QueryRow(ctx, "INSERT INTO products (name, stock, price) VALUES ('Mechanical Keyboard', 3, 120.00) RETURNING id").Scan(&prodID)
	_ = conn.QueryRow(ctx, "INSERT INTO store_customers (name) VALUES ('Judy') RETURNING id").Scan(&custID)

	// Try ordering 5 items (Stock is only 3 -> Should fail and ROLLBACK)
	fmt.Println("Attempting to order 5 units (Stock = 3)...")
	err = placeOrder(ctx, conn, custID, prodID, 5)
	if err != nil {
		fmt.Printf("Transaction Rolled Back as Expected: %v\n", err)
	}

	// Try ordering 2 items (Stock is 3 -> Should succeed)
	fmt.Println("\nAttempting to order 2 units...")
	err = placeOrder(ctx, conn, custID, prodID, 2)
	if err != nil {
		log.Fatalf("Order failed: %v\n", err)
	}
	fmt.Println("Order placed successfully and stock updated!")
}
