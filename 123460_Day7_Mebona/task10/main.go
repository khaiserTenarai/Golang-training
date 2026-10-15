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
		fmt.Println("  ORDER & INVENTORY SYSTEM")
		fmt.Println("==============================")
		fmt.Println("1. Create Product (with stock)")
		fmt.Println("2. View Inventory")
		fmt.Println("3. Place Order (Auto-reduce stock)")
		fmt.Println("4. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}

		switch choice {
		case 1:
			createProduct(conn)
		case 2:
			viewInventory(conn)
		case 3:
			placeOrderWithInventoryCheck(conn)
		case 4:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func createProduct(conn *pgx.Conn) {
	fmt.Println("\n----- CREATE PRODUCT -----")
	name := readInput("Product Name: ")
	
	priceStr := readInput("Price: ")
	price, _ := strconv.ParseFloat(priceStr, 64)

	stockStr := readInput("Initial Stock Quantity: ")
	stock, _ := strconv.Atoi(stockStr)

	var id int
	err := conn.QueryRow(
		context.Background(),
		`INSERT INTO products (name, price, stock_quantity) VALUES ($1, $2, $3) RETURNING id`,
		name, price, stock,
	).Scan(&id)

	if err != nil {
		fmt.Println("Error creating product:", err)
	} else {
		fmt.Printf("Product created successfully with ID: %d\n", id)
	}
}

func viewInventory(conn *pgx.Conn) {
	fmt.Println("\n----- CURRENT INVENTORY -----")
	rows, err := conn.Query(context.Background(), `SELECT id, name, price, stock_quantity FROM products ORDER BY id`)
	if err != nil {
		fmt.Println("Error reading products:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, stock int
		var name string
		var price float64
		if err := rows.Scan(&id, &name, &price, &stock); err == nil {
			fmt.Printf("ID: %d | Name: %-15s | Price: $%-7.2f | Stock: %d\n", id, name, price, stock)
		}
	}
}

func placeOrderWithInventoryCheck(conn *pgx.Conn) {
	fmt.Println("\n----- PLACE ORDER -----")
	customerName := readInput("Enter Customer Name: ")

	ctx := context.Background()

	// 1. Begin Transaction
	tx, err := conn.Begin(ctx)
	if err != nil {
		fmt.Println("Failed to start transaction:", err)
		return
	}
	
	// Defer rollback to ensure database state is reverted if the function exits due to an error.
	defer tx.Rollback(ctx)

	// 2. Create the main order record
	var orderID int
	err = tx.QueryRow(ctx, `INSERT INTO orders (customer_name) VALUES ($1) RETURNING id`, customerName).Scan(&orderID)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return
	}

	productIDStr := readInput("Enter Product ID to purchase: ")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		fmt.Println("Invalid Product ID.")
		return
	}

	qtyStr := readInput("Enter Quantity: ")
	requestedQty, err := strconv.Atoi(qtyStr)
	if err != nil || requestedQty <= 0 {
		fmt.Println("Invalid quantity.")
		return
	}

	// 3. Check current stock and lock the row for update to prevent race conditions
	var currentStock int
	var price float64
	err = tx.QueryRow(ctx, `SELECT stock_quantity, price FROM products WHERE id = $1 FOR UPDATE`, productID).Scan(&currentStock, &price)
	if err != nil {
		fmt.Println("Error: Product not found.")
		return // Triggers rollback
	}

	// 4. Validate stock sufficiency
	if currentStock < requestedQty {
		fmt.Printf("Transaction Failed: Insufficient stock! Only %d left in inventory. Rolling back order.\n", currentStock)
		return // Triggers rollback, completely canceling the order creation
	}

	// 5. Reduce the stock quantity
	_, err = tx.Exec(ctx, `UPDATE products SET stock_quantity = stock_quantity - $1 WHERE id = $2`, requestedQty, productID)
	if err != nil {
		fmt.Println("Error updating stock:", err)
		return
	}

	// 6. Insert the order item
	_, err = tx.Exec(
		ctx,
		`INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`,
		orderID, productID, requestedQty, price,
	)
	if err != nil {
		fmt.Println("Error adding item to order:", err)
		return
	}

	// 7. Commit Transaction (Everything succeeds)
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}

	fmt.Printf("Order #%d placed successfully! Stock for Product ID %d reduced by %d.\n", orderID, productID, requestedQty)
}