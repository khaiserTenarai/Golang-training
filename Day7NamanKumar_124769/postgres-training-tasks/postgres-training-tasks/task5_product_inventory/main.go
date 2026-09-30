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
	connString := "postgres://postgres:password@localhost:5432/inventorydb"
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
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
		fmt.Println("5. Increase Stock")
		fmt.Println("6. Decrease Stock")
		fmt.Println("7. Low Stock Search")
		fmt.Println("8. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			createProduct(ctx, db, reader)
		case "2":
			listProducts(ctx, db)
		case "3":
			updateProduct(ctx, db, reader)
		case "4":
			deleteProduct(ctx, db, reader)
		case "5":
			changeStock(ctx, db, reader, 1)
		case "6":
			changeStock(ctx, db, reader, -1)
		case "7":
			lowStockSearch(ctx, db, reader)
		case "8":
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

	fmt.Print("Initial stock: ")
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

func updateProduct(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Product ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	fmt.Print("New name: ")
	name := readLine(reader)

	fmt.Print("New price: ")
	price, _ := strconv.ParseFloat(readLine(reader), 64)

	result, err := db.Exec(ctx, `UPDATE products SET name = $1, price = $2 WHERE id = $3`, name, price, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Product not found")
		return
	}
	fmt.Println("Product updated")
}

func deleteProduct(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Product ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	result, err := db.Exec(ctx, `DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Product not found")
		return
	}
	fmt.Println("Product deleted")
}

func changeStock(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader, sign int) {
	fmt.Print("Product ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Quantity: ")
	qty, _ := strconv.Atoi(readLine(reader))

	if sign < 0 {
		var current int
		if err := db.QueryRow(ctx, `SELECT stock FROM products WHERE id = $1`, id).Scan(&current); err != nil {
			fmt.Println("Error:", err)
			return
		}
		if current < qty {
			fmt.Println("Insufficient stock")
			return
		}
	}

	result, err := db.Exec(ctx, `UPDATE products SET stock = stock + $1 WHERE id = $2`, sign*qty, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Product not found")
		return
	}
	fmt.Println("Stock updated")
}

func lowStockSearch(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Threshold: ")
	threshold, _ := strconv.Atoi(readLine(reader))

	rows, err := db.Query(ctx, `SELECT id, name, stock FROM products WHERE stock < $1 ORDER BY stock`, threshold)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var id, stock int
		var name string
		if err := rows.Scan(&id, &name, &stock); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("ID: %d | Name: %s | Stock: %d\n", id, name, stock)
		found = true
	}
	if !found {
		fmt.Println("No low-stock products")
	}
}
