package repository

import (
	"database/sql"
	"task09_ecommerce_order_system/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(p models.Product) (int, error) {
	var id int
	err := r.DB.QueryRow(`INSERT INTO ecom_products (name, price, stock) VALUES ($1, $2, $3) RETURNING id`,
		p.Name, p.Price, p.Stock).Scan(&id)
	return id, err
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	rows, err := r.DB.Query(`SELECT id, name, price, stock, created_at FROM ecom_products ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *ProductRepository) GetByID(id int) (models.Product, error) {
	var p models.Product
	err := r.DB.QueryRow(`SELECT id, name, price, stock, created_at FROM ecom_products WHERE id=$1`, id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt)
	return p, err
}
