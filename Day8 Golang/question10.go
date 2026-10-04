package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=Day7 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()

	err = placeOrder(ctx, db, 1, 5)
	if err != nil {
		log.Fatalf("Order failed: %v", err)
	}

	fmt.Println("Order placed successfully")
}

func placeOrder(ctx context.Context, db *sql.DB, productID int, quantity int) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var stock int
	var price float64
	err = tx.QueryRowContext(ctx, `
		SELECT stock_quantity, price 
		FROM products 
		WHERE id = $1 FOR UPDATE`, productID).Scan(&stock, &price)
	if err != nil {
		return err
	}

	if stock < quantity {
		return errors.New("insufficient stock")
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE products 
		SET stock_quantity = stock_quantity - $1 
		WHERE id = $2`, quantity, productID)
	if err != nil {
		return err
	}

	totalAmount := price * float64(quantity)
	_, err = tx.ExecContext(ctx, `
		INSERT INTO orders (product_id, quantity, total_amount) 
		VALUES ($1, $2, $3)`, productID, quantity, totalAmount)
	if err != nil {
		return err
	}

	return tx.Commit()
}