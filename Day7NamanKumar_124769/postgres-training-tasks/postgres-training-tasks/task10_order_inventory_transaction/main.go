package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connString := "postgres://postgres:password@localhost:5432/orderinventorydb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Connection failed:", err)
		os.Exit(1)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Println("\n1. Create Product")
		fmt.Println("2. List Products")
		fmt.Println("3. Place Order")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			createProduct(ctx, db, reader)
		case "2":
			listProducts(ctx, db)
		case "3":
			placeOrder(ctx, db, reader)
		case "4":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func createProduct(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Name: ")
	name := readLine(reader)

	fmt.Print("Stock: ")
	stock, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Price: ")
	price, _ := strconv.ParseFloat(readLine(reader), 64)

	var id int
	err := db.QueryRow(ctx, `INSERT INTO products (name, stock, price) VALUES ($1, $2, $3) RETURNING id`, name, stock, price).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created product with ID:", id)
}

func listProducts(ctx context.Context, db *pgxpool.Pool) {
	rows, err := db.Query(ctx, `SELECT id, name, stock, price FROM products ORDER BY id`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, stock int
		var name string
		var price float64
		if err := rows.Scan(&id, &name, &stock, &price); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("ID: %d | Name: %s | Stock: %d | Price: %.2f\n", id, name, stock, price)
	}
}

func placeOrder(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Product ID: ")
	productID, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Quantity: ")
	qty, _ := strconv.Atoi(readLine(reader))

	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer tx.Rollback(ctx)

	var stock int
	err = tx.QueryRow(ctx, `SELECT stock FROM products WHERE id = $1 FOR UPDATE`, productID).Scan(&stock)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if stock < qty {
		fmt.Println("Insufficient stock, order cancelled")
		return
	}

	var orderID int
	err = tx.QueryRow(ctx, `INSERT INTO orders DEFAULT VALUES RETURNING id`).Scan(&orderID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = tx.Exec(ctx, `INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`, orderID, productID, qty)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = tx.Exec(ctx, `UPDATE products SET stock = stock - $1 WHERE id = $2`, qty, productID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Order placed with ID:", orderID)
}
