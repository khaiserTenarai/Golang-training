package repository

import (
	"database/sql"
	"task05_product_inventory/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(p models.Product) (int, error) {
	var id int
	err := r.DB.QueryRow(
		`INSERT INTO products (name, category, price, stock) VALUES ($1, $2, $3, $4) RETURNING id`,
		p.Name, p.Category, p.Price, p.Stock).Scan(&id)
	return id, err
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	rows, err := r.DB.Query(`SELECT id, name, category, price, stock, created_at FROM products ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *ProductRepository) GetByID(id int) (models.Product, error) {
	var p models.Product
	err := r.DB.QueryRow(`SELECT id, name, category, price, stock, created_at FROM products WHERE id=$1`, id).
		Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Stock, &p.CreatedAt)
	return p, err
}

func (r *ProductRepository) Update(p models.Product) error {
	_, err := r.DB.Exec(`UPDATE products SET name=$1, category=$2, price=$3, stock=$4 WHERE id=$5`,
		p.Name, p.Category, p.Price, p.Stock, p.ID)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM products WHERE id=$1`, id)
	return err
}

func (r *ProductRepository) IncreaseStock(id, qty int) error {
	_, err := r.DB.Exec(`UPDATE products SET stock = stock + $1 WHERE id = $2`, qty, id)
	return err
}

func (r *ProductRepository) DecreaseStock(id, qty int) error {
	_, err := r.DB.Exec(`UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1`, qty, id)
	return err
}

func (r *ProductRepository) GetLowStock(threshold int) ([]models.Product, error) {
	rows, err := r.DB.Query(
		`SELECT id, name, category, price, stock, created_at FROM products WHERE stock <= $1 ORDER BY stock ASC`, threshold)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *ProductRepository) SearchByName(name string) ([]models.Product, error) {
	rows, err := r.DB.Query(
		`SELECT id, name, category, price, stock, created_at FROM products WHERE name ILIKE $1 ORDER BY id`,
		"%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}
