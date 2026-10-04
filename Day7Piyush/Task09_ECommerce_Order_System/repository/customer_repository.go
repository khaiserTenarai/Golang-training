package repository

import (
	"database/sql"
	"task09_ecommerce_order_system/models"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

func (r *CustomerRepository) Create(c models.Customer) (int, error) {
	var id int
	err := r.DB.QueryRow(`INSERT INTO ecom_customers (name, email) VALUES ($1, $2) RETURNING id`,
		c.Name, c.Email).Scan(&id)
	return id, err
}

func (r *CustomerRepository) GetAll() ([]models.Customer, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, created_at FROM ecom_customers ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.CreatedAt); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, rows.Err()
}
