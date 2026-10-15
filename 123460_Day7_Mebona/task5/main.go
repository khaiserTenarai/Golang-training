package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

var reader = bufio.NewReader(os.Stdin)

func readInput(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	connString := "postgres://postgres:pgadmin@localhost:5432/employee_db"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL connected successfully")

	for {
		fmt.Println("\n==============================")
		fmt.Println("     PRODUCT INVENTORY CLI")
		fmt.Println("==============================")
		fmt.Println("1. Create Product")
		fmt.Println("2. List All Products")
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
		fmt.Println("5. Increase Stock")
		fmt.Println("6. Decrease Stock")
		fmt.Println("7. Search Low-Stock Products")
		fmt.Println("8. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 8.")
			continue
		}

		switch choice {
		case 1:
			createProduct(conn)
		case 2:
			listProducts(conn)
		case 3:
			updateProduct(conn)
		case 4:
			deleteProduct(conn)
		case 5:
			adjustStock(conn, "increase")
		case 6:
			adjustStock(conn, "decrease")
		case 7:
			searchLowStock(conn)
		case 8:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select between 1 and 8.")
		}
	}
}

func createProduct(conn *pgx.Conn) {
	fmt.Println("\n----- CREATE PRODUCT -----")
	name := readInput("Product Name: ")
	category := readInput("Category: ")

	priceStr := readInput("Price: ")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Invalid price.")
		return
	}

	stockStr := readInput("Stock Quantity: ")
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		fmt.Println("Invalid stock quantity.")
		return
	}

	thresholdStr := readInput("Low Stock Threshold (default 5): ")
	threshold := 5
	if thresholdStr != "" {
		if t, err := strconv.Atoi(thresholdStr); err == nil {
			threshold = t
		}
	}

	var id int
	err = conn.QueryRow(
		context.Background(),
		`INSERT INTO products (name, category, price, stock_quantity, min_threshold) 
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		name, category, price, stock, threshold,
	).Scan(&id)

	if err != nil {
		fmt.Println("Error creating product:", err)
	} else {
		fmt.Printf("Product created successfully with ID: %d\n", id)
	}
}

func listProducts(conn *pgx.Conn) {
	fmt.Println("\n----- ALL PRODUCTS -----")
	rows, err := conn.Query(
		context.Background(),
		`SELECT id, name, category, price, stock_quantity, min_threshold FROM products ORDER BY id`,
	)
	if err != nil {
		fmt.Println("Error reading products:", err)
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var id, stock, threshold int
		var name, category string
		var price float64

		if err := rows.Scan(&id, &name, &category, &price, &stock, &threshold); err != nil {
			continue
		}
		count++
		fmt.Printf("ID: %d | Name: %s | Category: %s | Price: $%.2f | Stock: %d | Min Threshold: %d\n", 
			id, name, category, price, stock, threshold)
	}

	if count == 0 {
		fmt.Println("No products found in inventory.")
	}
}

func updateProduct(conn *pgx.Conn) {
	fmt.Println("\n----- UPDATE PRODUCT -----")
	idStr := readInput("Enter Product ID to update: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	name := readInput("New Product Name: ")
	category := readInput("New Category: ")

	priceStr := readInput("New Price: ")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Invalid price.")
		return
	}

	result, err := conn.Exec(
		context.Background(),
		`UPDATE products SET name = $1, category = $2, price = $3 WHERE id = $4`,
		name, category, price, id,
	)
	if err != nil {
		fmt.Println("Error updating product:", err)
	} else if result.RowsAffected() == 0 {
		fmt.Println("No product found with that ID.")
	} else {
		fmt.Println("Product updated successfully.")
	}
}

func deleteProduct(conn *pgx.Conn) {
	fmt.Println("\n----- DELETE PRODUCT -----")
	idStr := readInput("Enter Product ID to delete: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	result, err := conn.Exec(context.Background(), `DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error deleting product:", err)
	} else if result.RowsAffected() == 0 {
		fmt.Println("No product found with that ID.")
	} else {
		fmt.Println("Product deleted successfully.")
	}
}

func adjustStock(conn *pgx.Conn, action string) {
	fmt.Printf("\n----- %s STOCK -----\n", strings.ToUpper(action))
	idStr := readInput("Enter Product ID: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	qtyStr := readInput("Enter quantity to " + action + ": ")
	qty, err := strconv.Atoi(qtyStr)
	if err != nil || qty <= 0 {
		fmt.Println("Invalid quantity amount.")
		return
	}

	var query string
	if action == "increase" {
		query = `UPDATE products SET stock_quantity = stock_quantity + $1 WHERE id = $2`
	} else {
		// Ensure stock doesn't drop below zero via query guard
		query = `UPDATE products SET stock_quantity = stock_quantity - $1 WHERE id = $2 AND stock_quantity >= $1`
	}

	result, err := conn.Exec(context.Background(), query, qty, id)
	if err != nil {
		fmt.Println("Error updating stock:", err)
	} else if result.RowsAffected() == 0 {
		fmt.Println("Stock update failed. Either the product ID does not exist or there is insufficient stock.")
	} else {
		fmt.Printf("Stock successfully %d for product ID %d.\n", qty, id)
	}
}

func searchLowStock(conn *pgx.Conn) {
	fmt.Println("\n----- LOW-STOCK PRODUCTS ALERT -----")
	rows, err := conn.Query(
		context.Background(),
		`SELECT id, name, category, price, stock_quantity, min_threshold 
		 FROM products 
		 WHERE stock_quantity <= min_threshold 
		 ORDER BY stock_quantity ASC`,
	)
	if err != nil {
		fmt.Println("Error fetching low stock products:", err)
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var id, stock, threshold int
		var name, category string
		var price float64

		if err := rows.Scan(&id, &name, &category, &price, &stock, &threshold); err != nil {
			continue
		}
		count++
		fmt.Printf("[LOW STOCK] ID: %d | Name: %s | Category: %s | Current Stock: %d (Threshold: %d)\n", 
			id, name, category, stock, threshold)
	}

	if count == 0 {
		fmt.Println("All products have sufficient stock levels.")
	}
}