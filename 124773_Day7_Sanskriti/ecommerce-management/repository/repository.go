package repository

import (
	"context"

	"ecommerce-management/model"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		DB: db,
	}
}

// Add Product
func (r *Repository) AddProduct(product *model.Product) error {

	query := `
		INSERT INTO products (name, price, stock, low_stock_limit)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Stock,
		product.LowStockLimit,
	)

	return err
}

// Get All Products
func (r *Repository) GetProducts() ([]model.Product, error) {

	query := `
		SELECT id, name, price, stock, low_stock_limit
		FROM products
		ORDER BY id
	`

	rows, err := r.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {

		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.LowStockLimit,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// Get Product By ID
func (r *Repository) GetProduct(id int) (*model.Product, error) {

	query := `
		SELECT id, name, price, stock, low_stock_limit
		FROM products
		WHERE id = $1
	`

	var product model.Product

	err := r.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.LowStockLimit,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Update Product
func (r *Repository) UpdateProduct(product *model.Product) error {

	query := `
		UPDATE products
		SET name = $1,
			price = $2,
			stock = $3,
			low_stock_limit = $4
		WHERE id = $5
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Stock,
		product.LowStockLimit,
		product.ID,
	)

	return err
}

// Delete Product
func (r *Repository) DeleteProduct(id int) error {

	query := `
		DELETE FROM products
		WHERE id = $1
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}

// Increase Stock
func (r *Repository) IncreaseStock(id int, quantity int) error {

	query := `
		UPDATE products
		SET stock = stock + $1
		WHERE id = $2
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		quantity,
		id,
	)

	return err
}

// Decrease Stock
func (r *Repository) DecreaseStock(id int, quantity int) error {

	query := `
		UPDATE products
		SET stock = stock - $1
		WHERE id = $2
		AND stock >= $1
	`

	result, err := r.DB.Exec(
		context.Background(),
		query,
		quantity,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// Get Low Stock Products
func (r *Repository) GetLowStockProducts() ([]model.Product, error) {

	query := `
		SELECT id, name, price, stock, low_stock_limit
		FROM products
		WHERE stock <= low_stock_limit
		ORDER BY stock
	`

	rows, err := r.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {

		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.LowStockLimit,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
