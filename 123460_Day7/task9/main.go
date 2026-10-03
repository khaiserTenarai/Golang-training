package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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
		fmt.Println("   E-COMMERCE ORDER SYSTEM")
		fmt.Println("==============================")
		fmt.Println("1. Create Customer")
		fmt.Println("2. Create Product")
		fmt.Println("3. Place Order")
		fmt.Println("4. View Complete Order Details (JOIN)")
		fmt.Println("5. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			continue
		}

		switch choice {
		case 1:
			createCustomer(conn)
		case 2:
			createProduct(conn)
		case 3:
			placeOrder(conn)
		case 4:
			viewOrderDetails(conn)
		case 5:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select between 1 and 5.")
		}
	}
}

func createCustomer(conn *pgx.Conn) {
	fmt.Println("\n----- CREATE CUSTOMER -----")
	name := readInput("Name: ")
	email := readInput("Email: ")

	var id int
	err := conn.QueryRow(
		context.Background(),
		`INSERT INTO customers (name, email) VALUES ($1, $2) RETURNING id`,
		name, email,
	).Scan(&id)

	if err != nil {
		fmt.Println("Error creating customer:", err)
	} else {
		fmt.Printf("Customer created successfully with ID: %d\n", id)
	}
}

func createProduct(conn *pgx.Conn) {
	fmt.Println("\n----- CREATE PRODUCT -----")
	name := readInput("Product Name: ")
	
	priceStr := readInput("Price: ")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price <= 0 {
		fmt.Println("Invalid price.")
		return
	}

	var id int
	err = conn.QueryRow(
		context.Background(),
		`INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`,
		name, price,
	).Scan(&id)

	if err != nil {
		fmt.Println("Error creating product:", err)
	} else {
		fmt.Printf("Product created successfully with ID: %d\n", id)
	}
}

func placeOrder(conn *pgx.Conn) {
	fmt.Println("\n----- PLACE ORDER -----")
	
	customerIDStr := readInput("Enter Customer ID: ")
	customerID, err := strconv.Atoi(customerIDStr)
	if err != nil {
		fmt.Println("Invalid Customer ID.")
		return
	}

	ctx := context.Background()
	tx, err := conn.Begin(ctx)
	if err != nil {
		fmt.Println("Failed to start transaction:", err)
		return
	}
	defer tx.Rollback(ctx)

	// Verify customer exists
	var customerName string
	err = tx.QueryRow(ctx, `SELECT name FROM customers WHERE id = $1`, customerID).Scan(&customerName)
	if err != nil {
		fmt.Println("Error: Customer not found.")
		return
	}

	// Create the main order record
	var orderID int
	err = tx.QueryRow(ctx, `INSERT INTO orders (customer_id) VALUES ($1) RETURNING id`, customerID).Scan(&orderID)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return
	}

	fmt.Printf("Order #%d created for %s. Let's add items.\n", orderID, customerName)

	itemCount := 0
	for {
		productIDStr := readInput("\nEnter Product ID (or press Enter to finish): ")
		if productIDStr == "" {
			break
		}

		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			fmt.Println("Invalid Product ID. Try again.")
			continue
		}

		// Fetch the current price of the product
		var productPrice float64
		err = tx.QueryRow(ctx, `SELECT price FROM products WHERE id = $1`, productID).Scan(&productPrice)
		if err != nil {
			fmt.Println("Error: Product not found.")
			continue
		}

		qtyStr := readInput("Enter Quantity: ")
		qty, err := strconv.Atoi(qtyStr)
		if err != nil || qty <= 0 {
			fmt.Println("Invalid quantity. Must be greater than 0.")
			continue
		}

		// Insert into order_items
		_, err = tx.Exec(
			ctx,
			`INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`,
			orderID, productID, qty, productPrice,
		)
		if err != nil {
			fmt.Println("Error adding item to order:", err)
			continue
		}
		
		fmt.Println("Item added successfully!")
		itemCount++
	}

	if itemCount == 0 {
		fmt.Println("No items added. Cancelling order.")
		return // Triggers rollback automatically
	}

	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}
	fmt.Printf("\nOrder #%d finalized successfully with %d items!\n", orderID, itemCount)
}

func viewOrderDetails(conn *pgx.Conn) {
	fmt.Println("\n----- VIEW ORDER DETAILS -----")
	orderIDStr := readInput("Enter Order ID: ")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		fmt.Println("Invalid Order ID.")
		return
	}

	// 4-Table JOIN Query
	query := `
		SELECT 
			o.id AS order_id,
			o.order_date,
			c.name AS customer_name,
			c.email AS customer_email,
			p.name AS product_name,
			oi.quantity,
			oi.price AS unit_price,
			(oi.quantity * oi.price) AS total_price
		FROM orders o
		JOIN customers c ON o.customer_id = c.id
		JOIN order_items oi ON o.id = oi.order_id
		JOIN products p ON oi.product_id = p.id
		WHERE o.id = $1
	`

	rows, err := conn.Query(context.Background(), query, orderID)
	if err != nil {
		fmt.Println("Error fetching order details:", err)
		return
	}
	defer rows.Close()

	var orderDate time.Time
	var custName, custEmail string
	var grandTotal float64
	hasRows := false

	for rows.Next() {
		var oID, qty int
		var prodName string
		var unitPrice, lineTotal float64

		if !hasRows {
			fmt.Println("\n--- ORDER SUMMARY ---")
		}

		err := rows.Scan(&oID, &orderDate, &custName, &custEmail, &prodName, &qty, &unitPrice, &lineTotal)
		if err != nil {
			fmt.Println("Error reading row:", err)
			continue
		}

		if !hasRows {
			fmt.Printf("Order ID : %d\n", oID)
			fmt.Printf("Date     : %s\n", orderDate.Format("2006-01-02 15:04:05"))
			fmt.Printf("Customer : %s (%s)\n", custName, custEmail)
			fmt.Println("--------------------------------------------------")
			fmt.Printf("%-20s | %-5s | %-10s | %-10s\n", "Product", "Qty", "Unit Price", "Line Total")
			fmt.Println("--------------------------------------------------")
			hasRows = true
		}

		// The clean line without any invalid escapes
		fmt.Printf("%-20s | %-5d | $%-9.2f | $%-9.2f\n", prodName, qty, unitPrice, lineTotal)
		grandTotal += lineTotal
	}

	if !hasRows {
		fmt.Println("No order found with that ID or the order has no items.")
	} else {
		fmt.Println("--------------------------------------------------")
		fmt.Printf("GRAND TOTAL: $%.2f\n", grandTotal)
	}
}