// Package repository has all the direct database access for customers.
package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"question6-customer-management/model"
)

// ErrCustomerNotFound is returned when a customer ID doesn't exist.
var ErrCustomerNotFound = errors.New("customer not found")

// CreateCustomer inserts a new customer and returns its generated ID.
func CreateCustomer(db *sql.DB, c model.Customer) (int, error) {
	var id int
	query := `
		INSERT INTO customer (name, email, phone, address)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	err := db.QueryRow(query, c.Name, c.Email, c.Phone, c.Address).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create customer failed: %w", err)
	}
	return id, nil
}

// GetCustomer fetches one customer by ID.
func GetCustomer(db *sql.DB, id int) (model.Customer, error) {
	var c model.Customer
	query := `SELECT id, name, email, phone, address, created_at FROM customer WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.Address, &c.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Customer{}, fmt.Errorf("get customer failed: %w", ErrCustomerNotFound)
		}
		return model.Customer{}, fmt.Errorf("get customer failed: %w", err)
	}
	return c, nil
}

// UpdateCustomer updates an existing customer's details.
func UpdateCustomer(db *sql.DB, id int, name, email, phone, address string) error {
	query := `UPDATE customer SET name = $1, email = $2, phone = $3, address = $4 WHERE id = $5`
	result, err := db.Exec(query, name, email, phone, address, id)
	if err != nil {
		return fmt.Errorf("update customer failed: %w", err)
	}
	return checkRowsAffected(result, ErrCustomerNotFound)
}

// DeleteCustomer removes a customer by ID.
func DeleteCustomer(db *sql.DB, id int) error {
	result, err := db.Exec(`DELETE FROM customer WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete customer failed: %w", err)
	}
	return checkRowsAffected(result, ErrCustomerNotFound)
}

// CustomerSearchParams holds the optional filters and pagination
// options for SearchCustomers. Leave a field empty/zero to skip it.
type CustomerSearchParams struct {
	Name     string // partial, case-insensitive match
	Email    string // partial, case-insensitive match
	Phone    string // partial match
	Page     int    // 1-based; defaults to 1
	PageSize int    // defaults to 10
}

// SearchCustomers returns one page of customers matching the given
// filters (ordered by name), plus the total number of matches
// (ignoring pagination) so the caller knows how many pages exist.
func SearchCustomers(db *sql.DB, params CustomerSearchParams) ([]model.Customer, int, error) {
	var conditions []string
	var args []interface{}
	argPos := 1

	if params.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", argPos))
		args = append(args, "%"+params.Name+"%")
		argPos++
	}
	if params.Email != "" {
		conditions = append(conditions, fmt.Sprintf("email ILIKE $%d", argPos))
		args = append(args, "%"+params.Email+"%")
		argPos++
	}
	if params.Phone != "" {
		conditions = append(conditions, fmt.Sprintf("phone ILIKE $%d", argPos))
		args = append(args, "%"+params.Phone+"%")
		argPos++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM customer %s", whereClause)
	var total int
	if err := db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count customers failed: %w", err)
	}

	page := params.Page
	if page < 1 {
		page = 1
	}
	pageSize := params.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query := fmt.Sprintf(`
		SELECT id, name, email, phone, address, created_at
		FROM customer
		%s
		ORDER BY name ASC
		LIMIT $%d OFFSET $%d`,
		whereClause, argPos, argPos+1)

	args = append(args, pageSize, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("search customers failed: %w", err)
	}
	defer rows.Close()

	var customers []model.Customer
	for rows.Next() {
		var c model.Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.Address, &c.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("search customers failed: %w", err)
		}
		customers = append(customers, c)
	}
	return customers, total, rows.Err()
}

func checkRowsAffected(result sql.Result, notFoundErr error) error {
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check affected rows: %w", err)
	}
	if rows == 0 {
		return notFoundErr
	}
	return nil
}
