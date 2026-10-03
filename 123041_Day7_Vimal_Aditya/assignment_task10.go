package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {

	connString := "postgres://postgres:Info%40131@localhost:5432/go_training"

	conn, err := pgx.Connect(
		context.Background(),
		connString,
	)

	if err != nil {

		log.Fatal("Connection failed:", err)

	}

	defer conn.Close(context.Background())

	ctx := context.Background()

	fmt.Println("PostgreSQL connected successfully")

	// =====================================
	// ORDER DETAILS
	// =====================================

	productID := 1
	quantity := 2
	customerName := "Vimal Aditya"

	// =====================================
	// START TRANSACTION
	// =====================================

	tx, err := conn.Begin(ctx)

	if err != nil {

		log.Fatal(err)

	}

	// Rollback if error

	defer tx.Rollback(ctx)

	// =====================================
	// CHECK STOCK
	// =====================================

	var stock int

	err = tx.QueryRow(
		ctx,
		`
		SELECT stock

		FROM products

		WHERE id=$1

		FOR UPDATE
		`,
		productID,
	).Scan(&stock)

	if err != nil {

		log.Println("Product not found")
		return

	}

	fmt.Println("Available stock:", stock)

	if stock < quantity {

		fmt.Println("Insufficient stock")

		return

	}

	// =====================================
	// CREATE ORDER
	// =====================================

	var orderID int

	err = tx.QueryRow(
		ctx,
		`
		INSERT INTO orders
		(customer_name)

		VALUES($1)

		RETURNING id
		`,
		customerName,
	).Scan(&orderID)

	if err != nil {

		log.Println("Order creation failed")
		return

	}

	// =====================================
	// REDUCE STOCK
	// =====================================

	_, err = tx.Exec(
		ctx,
		`
		UPDATE products

		SET stock = stock - $1

		WHERE id=$2
		`,
		quantity,
		productID,
	)

	if err != nil {

		log.Println("Stock update failed")
		return

	}

	// =====================================
	// INSERT ORDER ITEM
	// =====================================

	_, err = tx.Exec(
		ctx,
		`
		INSERT INTO order_items
		(order_id,product_id,quantity)

		VALUES($1,$2,$3)
		`,
		orderID,
		productID,
		quantity,
	)

	if err != nil {

		log.Println("Order item failed")
		return

	}

	// =====================================
	// COMMIT
	// =====================================

	err = tx.Commit(ctx)

	if err != nil {

		log.Println("Commit failed")

	} else {

		fmt.Println("Order created successfully")

	}

	// =====================================
	// DISPLAY STOCK
	// =====================================

	var finalStock int

	err = conn.QueryRow(
		ctx,
		`
		SELECT stock

		FROM products

		WHERE id=$1
		`,
		productID,
	).Scan(&finalStock)

	if err == nil {

		fmt.Println("Remaining stock:", finalStock)

	}

}
