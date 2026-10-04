package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type CustomerFilter struct {
	Keyword  string
	Page     int
	PageSize int
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=gotraining sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	customerID := createCustomer(db, "John", "Doe", "john.doe@example.com", "555-0100")
	fmt.Printf("Created Customer ID: %d\n", customerID)

	c := getCustomer(db, customerID)
	fmt.Printf("Read Customer: %+v\n", c)

	updateCustomer(db, customerID, "John", "Smith", "john.smith@example.com", "555-0199")

	filter := CustomerFilter{
		Keyword:  "John",
		Page:     1,
		PageSize: 10,
	}
	
	results := searchCustomers(db, filter)
	for _, res := range results {
		fmt.Printf("Search Result: %+v\n", res)
	}

	deleteCustomer(db, customerID)
}

func createCustomer(db *sql.DB, firstName, lastName, email, phone string) int {
	var id int
	err := db.QueryRow(`
		INSERT INTO customers (first_name, last_name, email, phone) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		firstName, lastName, email, phone).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func getCustomer(db *sql.DB, id int) Customer {
	var c Customer
	err := db.QueryRow(`
		SELECT id, first_name, last_name, email, phone 
		FROM customers WHERE id = $1`, id).
		Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func updateCustomer(db *sql.DB, id int, firstName, lastName, email, phone string) {
	_, err := db.Exec(`
		UPDATE customers SET first_name = $1, last_name = $2, email = $3, phone = $4 
		WHERE id = $5`,
		firstName, lastName, email, phone, id)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteCustomer(db *sql.DB, id int) {
	_, err := db.Exec(`DELETE FROM customers WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
}

func searchCustomers(db *sql.DB, filter CustomerFilter) []Customer {
	query := `
		SELECT id, first_name, last_name, email, phone 
		FROM customers 
		WHERE 1=1
	`
	args := []interface{}{}
	argId := 1

	if filter.Keyword != "" {
		query += fmt.Sprintf(" AND (first_name ILIKE $%d OR last_name ILIKE $%d OR email ILIKE $%d)", argId, argId, argId)
		args = append(args, "%"+filter.Keyword+"%")
		argId++
	}

	pageSize := filter.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	page := filter.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query += fmt.Sprintf(" ORDER BY id ASC LIMIT $%d OFFSET $%d", argId, argId+1)
	args = append(args, pageSize, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone); err != nil {
			log.Fatal(err)
		}
		customers = append(customers, c)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return customers
}