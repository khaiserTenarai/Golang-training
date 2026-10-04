package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"task06_customer_management/models"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

func (r *CustomerRepository) Create(c models.Customer) (int, error) {
	var id int
	err := r.DB.QueryRow(
		`INSERT INTO customers (name, email, phone, city) VALUES ($1, $2, $3, $4) RETURNING id`,
		c.Name, c.Email, c.Phone, c.City).Scan(&id)
	return id, err
}

func (r *CustomerRepository) GetAll(page, pageSize int) ([]models.Customer, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}

	var total int
	r.DB.QueryRow(`SELECT COUNT(*) FROM customers`).Scan(&total)

	offset := (page - 1) * pageSize
	rows, err := r.DB.Query(
		`SELECT id, name, email, phone, city, created_at FROM customers ORDER BY id LIMIT $1 OFFSET $2`,
		pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.City, &c.CreatedAt); err != nil {
			return nil, 0, err
		}
		customers = append(customers, c)
	}
	return customers, total, rows.Err()
}

func (r *CustomerRepository) GetByID(id int) (models.Customer, error) {
	var c models.Customer
	err := r.DB.QueryRow(
		`SELECT id, name, email, phone, city, created_at FROM customers WHERE id=$1`, id).
		Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.City, &c.CreatedAt)
	return c, err
}

func (r *CustomerRepository) Update(c models.Customer) error {
	_, err := r.DB.Exec(`UPDATE customers SET name=$1, email=$2, phone=$3, city=$4 WHERE id=$5`,
		c.Name, c.Email, c.Phone, c.City, c.ID)
	return err
}

func (r *CustomerRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM customers WHERE id=$1`, id)
	return err
}

func (r *CustomerRepository) Search(keyword string) ([]models.Customer, error) {
	search := "%" + keyword + "%"
	conditions := []string{"name ILIKE $1", "email ILIKE $1", "phone ILIKE $1", "city ILIKE $1"}
	query := fmt.Sprintf(
		`SELECT id, name, email, phone, city, created_at FROM customers WHERE %s ORDER BY id`,
		strings.Join(conditions, " OR "))

	rows, err := r.DB.Query(query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.City, &c.CreatedAt); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, rows.Err()
}
