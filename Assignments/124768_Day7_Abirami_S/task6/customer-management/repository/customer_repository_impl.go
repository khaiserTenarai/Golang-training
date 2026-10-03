package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"customer-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrCustomerNotFound = errors.New("customer not found")

type CustomerRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) CustomerRepository {
	return &CustomerRepositoryImpl{
		db: db,
	}
}

func (r *CustomerRepositoryImpl) CreateCustomer(customer model.Customer) error {
	fmt.Println("\n----- CREATE CUSTOMER -----")

	query := `
		INSERT INTO customers (name, email, phone, city)
		VALUES ($1, $2, $3, $4)
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		customer.Name,
		customer.Email,
		customer.Phone,
		customer.City,
	)

	if err != nil {
		return err
	}

	fmt.Println("Customer created successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

func (r *CustomerRepositoryImpl) GetCustomer(id int) (*model.Customer, error) {
	fmt.Println("\n----- READ CUSTOMER -----")

	var customer model.Customer

	query := `
		SELECT id, name, email, phone, city
		FROM customers
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.City,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepositoryImpl) GetAllCustomers() ([]model.Customer, error) {
	fmt.Println("\n----- READ ALL CUSTOMERS -----")

	query := `
		SELECT id, name, email, phone, city
		FROM customers
		ORDER BY id
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []model.Customer

	for rows.Next() {
		var customer model.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Phone,
			&customer.City,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepositoryImpl) UpdateCustomer(customer model.Customer) error {
	fmt.Println("\n----- UPDATE CUSTOMER -----")

	query := `
		UPDATE customers
		SET name = $1,
		    email = $2,
		    phone = $3,
		    city = $4
		WHERE id = $5
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		customer.Name,
		customer.Email,
		customer.Phone,
		customer.City,
		customer.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrCustomerNotFound
	}

	fmt.Println("Customer updated successfully")

	return nil
}

func (r *CustomerRepositoryImpl) DeleteCustomer(id int) error {
	fmt.Println("\n----- DELETE CUSTOMER -----")

	result, err := r.db.Exec(
		context.Background(),
		`DELETE FROM customers WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrCustomerNotFound
	}

	fmt.Println("Customer deleted successfully")

	return nil
}

func (r *CustomerRepositoryImpl) SearchCustomers(name string, email string, city string, page int, size int) ([]model.Customer, error) {
	fmt.Println("\n----- SEARCH CUSTOMERS -----")

	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 5
	}

	offset := (page - 1) * size

	query := `
		SELECT id, name, email, phone, city
		FROM customers
		WHERE 1 = 1
	`

	var args []interface{}
	argNumber := 1

	if strings.TrimSpace(name) != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argNumber)
		args = append(args, "%"+name+"%")
		argNumber++
	}

	if strings.TrimSpace(email) != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", argNumber)
		args = append(args, "%"+email+"%")
		argNumber++
	}

	if strings.TrimSpace(city) != "" {
		query += fmt.Sprintf(" AND city ILIKE $%d", argNumber)
		args = append(args, "%"+city+"%")
		argNumber++
	}

	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argNumber, argNumber+1)

	args = append(args, size, offset)

	rows, err := r.db.Query(
		context.Background(),
		query,
		args...,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []model.Customer

	for rows.Next() {
		var customer model.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Phone,
			&customer.City,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
