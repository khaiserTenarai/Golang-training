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
	connString := "postgres://postgres:password@localhost:5432/customerdb"
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
		fmt.Println("2. Update Customer")
		fmt.Println("3. Delete Customer")
		fmt.Println("4. Search Customers")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			createCustomer(ctx, db, reader)
		case "2":
			updateCustomer(ctx, db, reader)
		case "3":
			deleteCustomer(ctx, db, reader)
		case "4":
			searchCustomers(ctx, db, reader)
		case "5":
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
	fmt.Print("Name: ")
	name := readLine(reader)

	fmt.Print("Email: ")
	email := readLine(reader)

	fmt.Print("Phone: ")
	phone := readLine(reader)

	var id int
	err := db.QueryRow(ctx, `INSERT INTO customers (name, email, phone) VALUES ($1, $2, $3) RETURNING id`, name, email, phone).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created customer with ID:", id)
}

func updateCustomer(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Customer ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	fmt.Print("New name: ")
	name := readLine(reader)

	fmt.Print("New phone: ")
	phone := readLine(reader)

	result, err := db.Exec(ctx, `UPDATE customers SET name = $1, phone = $2 WHERE id = $3`, name, phone, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Customer not found")
		return
	}
	fmt.Println("Customer updated")
}

func deleteCustomer(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Customer ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	result, err := db.Exec(ctx, `DELETE FROM customers WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Customer not found")
		return
	}
	fmt.Println("Customer deleted")
}

func searchCustomers(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Name contains (blank to skip): ")
	name := readLine(reader)

	fmt.Print("Page number: ")
	page, _ := strconv.Atoi(readLine(reader))
	if page < 1 {
		page = 1
	}

	fmt.Print("Page size: ")
	pageSize, _ := strconv.Atoi(readLine(reader))
	if pageSize < 1 {
		pageSize = 10
	}

	query := "SELECT id, name, email, phone FROM customers WHERE 1=1"
	args := []interface{}{}
	argPos := 1

	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argPos)
		args = append(args, "%"+name+"%")
		argPos++
	}

	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, pageSize, (page-1)*pageSize)

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var id int
		var custName, email, phone string
		if err := rows.Scan(&id, &custName, &email, &phone); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("ID: %d | Name: %s | Email: %s | Phone: %s\n", id, custName, email, phone)
		found = true
	}
	if !found {
		fmt.Println("No matching customers")
	}
}
