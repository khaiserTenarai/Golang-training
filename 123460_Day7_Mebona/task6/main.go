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
	// Adjust connection string if necessary
	connString := "postgres://postgres:pgadmin@localhost:5432/employee_db"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL connected successfully")

	for {
		fmt.Println("\n==============================")
		fmt.Println("   CUSTOMER MANAGEMENT CLI")
		fmt.Println("==============================")
		fmt.Println("1. Create Customer")
		fmt.Println("2. List All Customers")
		fmt.Println("3. Update Customer")
		fmt.Println("4. Delete Customer")
		fmt.Println("5. Search Customers (with Pagination)")
		fmt.Println("6. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 6.")
			continue
		}

		switch choice {
		case 1:
			createCustomer(conn)
		case 2:
			listCustomers(conn)
		case 3:
			updateCustomer(conn)
		case 4:
			deleteCustomer(conn)
		case 5:
			searchCustomers(conn)
		case 6:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select between 1 and 6.")
		}
	}
}

func createCustomer(conn *pgx.Conn) {
	fmt.Println("\n----- CREATE CUSTOMER -----")
	name := readInput("Name: ")
	email := readInput("Email: ")
	phone := readInput("Phone: ")

	var id int
	err := conn.QueryRow(
		context.Background(),
		`INSERT INTO customers (name, email, phone) VALUES ($1, $2, $3) RETURNING id`,
		name, email, phone,
	).Scan(&id)

	if err != nil {
		fmt.Println("Error creating customer:", err)
	} else {
		fmt.Printf("Customer created successfully with ID: %d\n", id)
	}
}

func listCustomers(conn *pgx.Conn) {
	fmt.Println("\n----- ALL CUSTOMERS -----")
	rows, err := conn.Query(context.Background(), `SELECT id, name, email, phone FROM customers ORDER BY id`)
	if err != nil {
		fmt.Println("Error reading customers:", err)
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var id int
		var name, email, phone string
		if err := rows.Scan(&id, &name, &email, &phone); err != nil {
			continue
		}
		count++
		fmt.Printf("ID: %d | Name: %s | Email: %s | Phone: %s\n", id, name, email, phone)
	}

	if count == 0 {
		fmt.Println("No customers found.")
	}
}

func updateCustomer(conn *pgx.Conn) {
	fmt.Println("\n----- UPDATE CUSTOMER -----")
	idStr := readInput("Enter Customer ID to update: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	name := readInput("New Name: ")
	email := readInput("New Email: ")
	phone := readInput("New Phone: ")

	result, err := conn.Exec(
		context.Background(),
		`UPDATE customers SET name = $1, email = $2, phone = $3 WHERE id = $4`,
		name, email, phone, id,
	)
	if err != nil {
		fmt.Println("Error updating customer:", err)
	} else if result.RowsAffected() == 0 {
		fmt.Println("No customer found with that ID.")
	} else {
		fmt.Println("Customer updated successfully.")
	}
}

func deleteCustomer(conn *pgx.Conn) {
	fmt.Println("\n----- DELETE CUSTOMER -----")
	idStr := readInput("Enter Customer ID to delete: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	result, err := conn.Exec(context.Background(), `DELETE FROM customers WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error deleting customer:", err)
	} else if result.RowsAffected() == 0 {
		fmt.Println("No customer found with that ID.")
	} else {
		fmt.Println("Customer deleted successfully.")
	}
}

func searchCustomers(conn *pgx.Conn) {
	fmt.Println("\n----- SEARCH CUSTOMERS -----")
	
	searchTerm := readInput("Search by Name or Email (leave blank for all): ")
	
	limitStr := readInput("Items per page (limit, default 5): ")
	limit := 5
	if val, err := strconv.Atoi(limitStr); err == nil && val > 0 {
		limit = val
	}

	pageStr := readInput("Page number (default 1): ")
	page := 1
	if val, err := strconv.Atoi(pageStr); err == nil && val > 0 {
		page = val
	}
	offset := (page - 1) * limit

	// Prepare wildcard string in Go to avoid SQL syntax confusion
	searchPattern := "%" + searchTerm + "%"

	query := `
		SELECT id, name, email, phone 
		FROM customers 
		WHERE ($1 = '%%' OR name ILIKE $1 OR email ILIKE $1)
		ORDER BY id
		LIMIT $2 OFFSET $3
	`

	rows, err := conn.Query(context.Background(), query, searchPattern, limit, offset)
	if err != nil {
		fmt.Println("Search error:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("\n--- SEARCH RESULTS (Page %d) ---\n", page)
	count := 0
	for rows.Next() {
		var id int
		var name, email, phone string

		if err := rows.Scan(&id, &name, &email, &phone); err != nil {
			continue
		}
		count++
		fmt.Printf("ID: %d | Name: %s | Email: %s | Phone: %s\n", id, name, email, phone)
	}

	if count == 0 {
		fmt.Println("No customers found matching the criteria on this page.")
	}
}