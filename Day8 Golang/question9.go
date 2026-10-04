package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type OrderDetail struct {
	OrderID     int
	OrderDate   time.Time
	Customer    string
	Email       string
	ProductName string
	Quantity    int
	UnitPrice   float64
	LineTotal   float64
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=Day7 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	custID := createCustomer(db, "Alice Johnson", "alice@example.com")
	
	prodID1 := createProduct(db, "Laptop", 999.99)
	prodID2 := createProduct(db, "Wireless Mouse", 25.50)

	orderID := createOrder(db, custID, 1025.49)
	
	createOrderItem(db, orderID, prodID1, 1, 999.99)
	createOrderItem(db, orderID, prodID2, 1, 25.50)

	details := getOrderDetails(db, orderID)
	
	if len(details) > 0 {
		fmt.Printf("Order #%d | Date: %s\n", details[0].OrderID, details[0].OrderDate.Format(time.RFC3339))
		fmt.Printf("Customer: %s (%s)\n", details[0].Customer, details[0].Email)
		fmt.Println("Items:")
		for _, d := range details {
			fmt.Printf("- %s | Qty: %d | Unit Price: $%.2f | Line Total: $%.2f\n", d.ProductName, d.Quantity, d.UnitPrice, d.LineTotal)
		}
	}
}

func createCustomer(db *sql.DB, name, email string) int {
	var id int
	err := db.QueryRow(`
		INSERT INTO customers (name, email) 
		VALUES ($1, $2) RETURNING id`, name, email).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func createProduct(db *sql.DB, name string, price float64) int {
	var id int
	err := db.QueryRow(`
		INSERT INTO products (name, price) 
		VALUES ($1, $2) RETURNING id`, name, price).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func createOrder(db *sql.DB, customerID int, totalAmount float64) int {
	var id int
	err := db.QueryRow(`
		INSERT INTO orders (customer_id, total_amount) 
		VALUES ($1, $2) RETURNING id`, customerID, totalAmount).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func createOrderItem(db *sql.DB, orderID, productID, quantity int, unitPrice float64) {
	_, err := db.Exec(`
		INSERT INTO order_items (order_id, product_id, quantity, unit_price) 
		VALUES ($1, $2, $3, $4)`, orderID, productID, quantity, unitPrice)
	if err != nil {
		log.Fatal(err)
	}
}

func getOrderDetails(db *sql.DB, orderID int) []OrderDetail {
	query := `
		SELECT 
			o.id, 
			o.order_date, 
			c.name as customer_name, 
			c.email, 
			p.name as product_name, 
			oi.quantity, 
			oi.unit_price,
			(oi.quantity * oi.unit_price) as line_total
		FROM orders o
		JOIN customers c ON o.customer_id = c.id
		JOIN order_items oi ON o.id = oi.order_id
		JOIN products p ON oi.product_id = p.id
		WHERE o.id = $1
	`
	
	rows, err := db.Query(query, orderID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var details []OrderDetail
	for rows.Next() {
		var d OrderDetail
		if err := rows.Scan(&d.OrderID, &d.OrderDate, &d.Customer, &d.Email, &d.ProductName, &d.Quantity, &d.UnitPrice, &d.LineTotal); err != nil {
			log.Fatal(err)
		}
		details = append(details, d)
	}
	
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	
	return details
}