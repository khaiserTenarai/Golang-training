// Customer Management - an interactive CLI backed by PostgreSQL.
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

	"question6-customer-management/db"
	"question6-customer-management/model"
	"question6-customer-management/repository"
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

// readOptionalInt is like readInt, but an empty answer is allowed and
// returns `fallback` - used for pagination prompts where the user
// might just want the default.
func readOptionalInt(prompt string, fallback int) int {
	text := readLine(prompt)
	if text == "" {
		return fallback
	}
	value, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("That wasn't a number, using default:", fallback)
		return fallback
	}
	return value
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
		fmt.Println("\n--- Customer Management ---")
		fmt.Println("1. Add Customer")
		fmt.Println("2. View Customer by ID")
		fmt.Println("3. Update Customer")
		fmt.Println("4. Delete Customer")
		fmt.Println("5. Search Customers (with pagination)")
		fmt.Println("6. Exit")

		switch readLine("Enter your choice: ") {
		case "1":
			addCustomer(conn)
		case "2":
			viewCustomer(conn)
		case "3":
			updateCustomer(conn)
		case "4":
			deleteCustomer(conn)
		case "5":
			searchCustomers(conn)
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

func printCustomer(c model.Customer) {
	fmt.Printf("  [%d] %s | %s | %s | %s\n", c.ID, c.Name, c.Email, c.Phone, c.Address)
}

func addCustomer(conn *sql.DB) {
	name := readLine("Name: ")
	email := readLine("Email: ")
	phone := readLine("Phone: ")
	address := readLine("Address: ")

	c := model.Customer{Name: name, Email: email, Phone: phone, Address: address}
	id, err := repository.CreateCustomer(conn, c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer created with ID:", id)
}

func viewCustomer(conn *sql.DB) {
	id := readInt("Enter customer ID: ")
	c, err := repository.GetCustomer(conn, id)
	if err != nil {
		if errors.Is(err, repository.ErrCustomerNotFound) {
			fmt.Println("No customer found with that ID.")
			return
		}
		fmt.Println("Error:", err)
		return
	}
	printCustomer(c)
}

func updateCustomer(conn *sql.DB) {
	id := readInt("Enter the ID of the customer to update: ")
	name := readLine("New name: ")
	email := readLine("New email: ")
	phone := readLine("New phone: ")
	address := readLine("New address: ")

	err := repository.UpdateCustomer(conn, id, name, email, phone, address)
	if err != nil {
		if errors.Is(err, repository.ErrCustomerNotFound) {
			fmt.Println("No customer found with that ID.")
			return
		}
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer updated.")
}

func deleteCustomer(conn *sql.DB) {
	id := readInt("Enter the ID of the customer to delete: ")
	err := repository.DeleteCustomer(conn, id)
	if err != nil {
		if errors.Is(err, repository.ErrCustomerNotFound) {
			fmt.Println("No customer found with that ID.")
			return
		}
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer deleted.")
}

func searchCustomers(conn *sql.DB) {
	fmt.Println("Leave any of these blank to skip that filter.")
	name := readLine("Search by name contains: ")
	email := readLine("Search by email contains: ")
	phone := readLine("Search by phone contains: ")
	page := readOptionalInt("Page number (default 1): ", 1)
	pageSize := readOptionalInt("Results per page (default 10): ", 10)

	customers, total, err := repository.SearchCustomers(conn, repository.CustomerSearchParams{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Total matches:", total)
	if len(customers) == 0 {
		fmt.Println("No customers found on this page.")
		return
	}
	for _, c := range customers {
		printCustomer(c)
	}
}
