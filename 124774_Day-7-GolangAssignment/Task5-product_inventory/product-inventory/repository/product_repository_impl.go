package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"prodcut-inventory/model"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewProductRepository(
	db *pgxpool.Pool,
) ProductRepository {

	return &ProductRepositoryImpl{
		db: db,
	}
}

// CREATE
func (r *ProductRepositoryImpl) Save(
	product model.Product,
) error {

	query := `
		INSERT INTO products
			(name, description, price, stock, low_stock_limit)
		VALUES
			($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.LowStockLimit,
	)

	return err
}

// READ ONE
func (r *ProductRepositoryImpl) FindByID(
	id int,
) (model.Product, error) {

	var product model.Product

	query := `
		SELECT
			id,
			name,
			description,
			price,
			stock,
			low_stock_limit
		FROM products
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.LowStockLimit,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return product, ErrProductNotFound
		}

		return product, err
	}

	return product, nil
}

// READ ALL
func (r *ProductRepositoryImpl) FindAll() []model.Product {

	query := `
		SELECT
			id,
			name,
			description,
			price,
			stock,
			low_stock_limit
		FROM products
		ORDER BY id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return []model.Product{}
	}

	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {

		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.LowStockLimit,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return []model.Product{}
		}

		products = append(products, product)
	}

	return products
}

// UPDATE
func (r *ProductRepositoryImpl) Update(
	product model.Product,
) error {

	query := `
		UPDATE products
		SET
			name = $1,
			description = $2,
			price = $3,
			stock = $4,
			low_stock_limit = $5
		WHERE id = $6
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.LowStockLimit,
		product.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	return nil
}

// DELETE
func (r *ProductRepositoryImpl) Delete(
	id int,
) error {

	result, err := r.db.Exec(
		context.Background(),
		"DELETE FROM products WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	return nil
}

// INCREASE STOCK
func (r *ProductRepositoryImpl) IncreaseStock(
	id int,
	quantity int,
) error {

	query := `
		UPDATE products
		SET stock = stock + $1
		WHERE id = $2
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		quantity,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	return nil
}

// DECREASE STOCK
func (r *ProductRepositoryImpl) DecreaseStock(
	id int,
	quantity int,
) error {

	query := `
		UPDATE products
		SET stock = stock - $1
		WHERE id = $2
		AND stock >= $1
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		quantity,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New(
			"product not found or insufficient stock",
		)
	}

	return nil
}

// LOW STOCK
func (r *ProductRepositoryImpl) FindLowStock() []model.Product {

	query := `
		SELECT
			id,
			name,
			description,
			price,
			stock,
			low_stock_limit
		FROM products
		WHERE stock <= low_stock_limit
		ORDER BY stock
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return []model.Product{}
	}

	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {

		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.LowStockLimit,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return []model.Product{}
		}

		products = append(products, product)
	}

	return products
}
