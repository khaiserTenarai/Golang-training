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
	connString := "postgres://postgres:password@localhost:5432/ecommercedb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Connection failed:", err)
		os.Exit(1)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Println("\n1. Create Customer")
		fmt.Println("2. Create Product")
		fmt.Println("3. Create Order")
		fmt.Println("4. Add Item to Order")
		fmt.Println("5. View Order Details")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			createCustomer(ctx, db, reader)
		case "2":
			createProduct(ctx, db, reader)
		case "3":
			createOrder(ctx, db, reader)
		case "4":
			addOrderItem(ctx, db, reader)
		case "5":
			viewOrderDetails(ctx, db, reader)
		case "6":
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

func createCustomer(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Customer name: ")
	name := readLine(reader)

	var id int
	err := db.QueryRow(ctx, `INSERT INTO customers (name) VALUES ($1) RETURNING id`, name).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created customer with ID:", id)
}

func createProduct(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Product name: ")
	name := readLine(reader)

	fmt.Print("Price: ")
	price, _ := strconv.ParseFloat(readLine(reader), 64)

	var id int
	err := db.QueryRow(ctx, `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`, name, price).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created product with ID:", id)
}

func createOrder(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Customer ID: ")
	custID, _ := strconv.Atoi(readLine(reader))

	var id int
	err := db.QueryRow(ctx, `INSERT INTO orders (customer_id) VALUES ($1) RETURNING id`, custID).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created order with ID:", id)
}

func addOrderItem(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Order ID: ")
	orderID, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Product ID: ")
	productID, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Quantity: ")
	qty, _ := strconv.Atoi(readLine(reader))

	_, err := db.Exec(ctx, `INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`, orderID, productID, qty)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Item added to order")
}

func viewOrderDetails(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Order ID: ")
	orderID, _ := strconv.Atoi(readLine(reader))

	query := `
		SELECT o.id, c.name, p.name, oi.quantity, p.price, (oi.quantity * p.price) AS line_total
		FROM orders o
		JOIN customers c ON o.customer_id = c.id
		JOIN order_items oi ON oi.order_id = o.id
		JOIN products p ON oi.product_id = p.id
		WHERE o.id = $1
	`
	rows, err := db.Query(ctx, query, orderID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	found := false
	var total float64
	for rows.Next() {
		var oID int
		var customerName, productName string
		var quantity int
		var price, lineTotal float64
		if err := rows.Scan(&oID, &customerName, &productName, &quantity, &price, &lineTotal); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Order #%d | Customer: %s | Product: %s | Qty: %d | Price: %.2f | Line Total: %.2f\n",
			oID, customerName, productName, quantity, price, lineTotal)
		total += lineTotal
		found = true
	}
	if !found {
		fmt.Println("Order not found or has no items")
		return
	}
	fmt.Printf("Order Total: %.2f\n", total)
}
