package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"customer-management/model"
)

type PostgresCustomerRepository struct {
	db *pgxpool.Pool
}

func NewPostgresCustomerRepository(db *pgxpool.Pool) *PostgresCustomerRepository {
	return &PostgresCustomerRepository{db: db}
}

func (r *PostgresCustomerRepository) AddCustomer(customer model.Customer) error {
	query := `
        INSERT INTO customers (name, email, phone, city)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.Exec(context.Background(), query,
		customer.Name, customer.Email, customer.Phone, customer.City)
	return err
}

func (r *PostgresCustomerRepository) GetCustomerByID(id int) (model.Customer, error) {
	query := `
        SELECT id, name, email, phone, city, created_at
        FROM customers
        WHERE id = $1
    `

	var customer model.Customer
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.City,
		&customer.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.Customer{}, fmt.Errorf("customer with ID %d not found", id)
	}

	return customer, err
}

func (r *PostgresCustomerRepository) GetAllCustomers(limit int, offset int) ([]model.Customer, error) {
	query := `
        SELECT id, name, email, phone, city, created_at
        FROM customers
        ORDER BY id
        LIMIT $1 OFFSET $2
    `

	rows, err := r.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Phone,
			&customer.City,
			&customer.CreatedAt,
		); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, rows.Err()
}

func (r *PostgresCustomerRepository) UpdateCustomer(customer model.Customer) error {
	query := `
        UPDATE customers
        SET name = $1, email = $2, phone = $3, city = $4
        WHERE id = $5
    `

	result, err := r.db.Exec(context.Background(), query,
		customer.Name, customer.Email, customer.Phone, customer.City, customer.ID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("customer with ID %d not found", customer.ID)
	}

	return nil
}

func (r *PostgresCustomerRepository) DeleteCustomer(id int) error {
	query := `DELETE FROM customers WHERE id = $1`

	result, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("customer with ID %d not found", id)
	}

	return nil
}

func (r *PostgresCustomerRepository) SearchCustomers(keyword string, limit int, offset int) ([]model.Customer, error) {
	query := `
        SELECT id, name, email, phone, city, created_at
        FROM customers
        WHERE name ILIKE $1
           OR email ILIKE $1
           OR phone ILIKE $1
           OR city ILIKE $1
        ORDER BY id
        LIMIT $2 OFFSET $3
    `

	searchTerm := "%" + keyword + "%"

	rows, err := r.db.Query(context.Background(), query, searchTerm, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Phone,
			&customer.City,
			&customer.CreatedAt,
		); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, rows.Err()
}

func (r *PostgresCustomerRepository) CountCustomers() (int, error) {
	query := `SELECT COUNT(*) FROM customers`

	var count int
	err := r.db.QueryRow(context.Background(), query).Scan(&count)
	return count, err
}
