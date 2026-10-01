// Product Inventory - an interactive CLI backed by PostgreSQL.
//
// Before running: apply sql/schema.sql to your database, then
//   go mod tidy
//   go run .
package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"question5-product-inventory/db"
	"question5-product-inventory/model"
	"question5-product-inventory/repository"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt(prompt string) int {
	for {
		text := readLine(prompt)
		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Please enter a whole number.")
			continue
		}
		return value
	}
}

func readFloat(prompt string) float64 {
	for {
		text := readLine(prompt)
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}
		return value
	}
}

func main() {
	conn, err := db.Connect()
	if err != nil {
		fmt.Println("Could not connect to the database:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to Postgres successfully.")

	for {
		fmt.Println("\n--- Product Inventory ---")
		fmt.Println("1. Add Product")
		fmt.Println("2. View Product by ID")
		fmt.Println("3. List All Products")
		fmt.Println("4. Update Product")
		fmt.Println("5. Delete Product")
		fmt.Println("6. Increase Stock")
		fmt.Println("7. Decrease Stock")
		fmt.Println("8. Search Low Stock Products")
		fmt.Println("9. Exit")

		switch readLine("Enter your choice: ") {
		case "1":
			addProduct(conn)
		case "2":
			viewProduct(conn)
		case "3":
			listProducts(conn)
		case "4":
			updateProduct(conn)
		case "5":
			deleteProduct(conn)
		case "6":
			increaseStock(conn)
		case "7":
			decreaseStock(conn)
		case "8":
			searchLowStock(conn)
		case "9":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

func printProduct(p model.Product) {
	fmt.Printf("  [%d] %s (SKU: %s) | Category: %s | Price: %.2f | Quantity: %d\n",
		p.ID, p.Name, p.SKU, p.Category, p.Price, p.Quantity)
}

func addProduct(conn *sql.DB) {
	name := readLine("Product name: ")
	sku := readLine("SKU (unique code): ")
	category := readLine("Category: ")
	price := readFloat("Price: ")
	quantity := readInt("Initial quantity: ")

	p := model.Product{Name: name, SKU: sku, Category: category, Price: price, Quantity: quantity}
	id, err := repository.CreateProduct(conn, p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Product created with ID:", id)
}

func viewProduct(conn *sql.DB) {
	id := readInt("Enter product ID: ")
	p, err := repository.GetProduct(conn, id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			fmt.Println("No product found with that ID.")
			return
		}
		fmt.Println("Error:", err)
		return
	}
	printProduct(p)
}

func listProducts(conn *sql.DB) {
	products, err := repository.ListProducts(conn)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(products) == 0 {
		fmt.Println("No products yet.")
		return
	}
	for _, p := range products {
		printProduct(p)
	}
}

func updateProduct(conn *sql.DB) {
	id := readInt("Enter the ID of the product to update: ")
	name := readLine("New name: ")
	category := readLine("New category: ")
	price := readFloat("New price: ")

	err := repository.UpdateProduct(conn, id, name, category, price)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			fmt.Println("No product found with that ID.")
			return
		}
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Product updated.")
}

func deleteProduct(conn *sql.DB) {
	id := readInt("Enter the ID of the product to delete: ")
	err := repository.DeleteProduct(conn, id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			fmt.Println("No product found with that ID.")
			return
		}
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Product deleted.")
}

func increaseStock(conn *sql.DB) {
	id := readInt("Enter product ID: ")
	amount := readInt("Amount to add to stock: ")

	err := repository.IncreaseStock(conn, id, amount)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			fmt.Println("No product found with that ID.")
			return
		}
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Stock increased.")
}

func decreaseStock(conn *sql.DB) {
	id := readInt("Enter product ID: ")
	amount := readInt("Amount to remove from stock: ")

	err := repository.DecreaseStock(conn, id, amount)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			fmt.Println("No product found with that ID.")
			return
		}
		if errors.Is(err, repository.ErrInsufficientStock) {
			fmt.Println("Not enough stock available for that decrease.")
			return
		}
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Stock decreased.")
}

func searchLowStock(conn *sql.DB) {
	threshold := readInt("Show products with quantity at or below: ")
	products, err := repository.SearchLowStock(conn, threshold)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(products) == 0 {
		fmt.Println("No products at or below that stock level.")
		return
	}
	fmt.Println("Low stock products:")
	for _, p := range products {
		printProduct(p)
	}
}
