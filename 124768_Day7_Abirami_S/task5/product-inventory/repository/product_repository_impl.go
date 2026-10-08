package repository

import (
	"context"
	"errors"
	"fmt"

	"product-inventory/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) CreateProduct(product model.Product) error {
	fmt.Println("\n----- CREATE PRODUCT -----")

	query := `
		INSERT INTO products (name, price, stock, low_stock_threshold)
		VALUES ($1, $2, $3, $4)
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Stock,
		product.LowStockThreshold,
	)

	if err != nil {
		return err
	}

	fmt.Println("Product created successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

func (r *ProductRepositoryImpl) GetProduct(id int) (*model.Product, error) {
	fmt.Println("\n----- READ PRODUCT -----")

	var product model.Product

	query := `
		SELECT id, name, price, stock, low_stock_threshold
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
		&product.Price,
		&product.Stock,
		&product.LowStockThreshold,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositoryImpl) GetAllProducts() ([]model.Product, error) {
	fmt.Println("\n----- READ ALL PRODUCTS -----")

	query := `
		SELECT id, name, price, stock, low_stock_threshold
		FROM products
		ORDER BY id
	`

	rows, err := r.db.Query(context.Background(), query)
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
			&product.LowStockThreshold,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(product model.Product) error {
	fmt.Println("\n----- UPDATE PRODUCT -----")

	query := `
		UPDATE products
		SET name = $1,
		    price = $2,
		    stock = $3,
		    low_stock_threshold = $4
		WHERE id = $5
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Stock,
		product.LowStockThreshold,
		product.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	fmt.Println("Product updated successfully")

	return nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id int) error {
	fmt.Println("\n----- DELETE PRODUCT -----")

	result, err := r.db.Exec(
		context.Background(),
		`DELETE FROM products WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	fmt.Println("Product deleted successfully")

	return nil
}

func (r *ProductRepositoryImpl) IncreaseStock(id int, quantity int) error {
	fmt.Println("\n----- INCREASE STOCK -----")

	result, err := r.db.Exec(
		context.Background(),
		`UPDATE products SET stock = stock + $1 WHERE id = $2`,
		quantity,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	fmt.Println("Stock increased successfully")

	return nil
}

func (r *ProductRepositoryImpl) DecreaseStock(id int, quantity int) error {
	fmt.Println("\n----- DECREASE STOCK -----")

	result, err := r.db.Exec(
		context.Background(),
		`UPDATE products
		 SET stock = stock - $1
		 WHERE id = $2 AND stock >= $1`,
		quantity,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("insufficient stock or product not found")
	}

	fmt.Println("Stock decreased successfully")

	return nil
}

func (r *ProductRepositoryImpl) GetLowStockProducts() ([]model.Product, error) {
	fmt.Println("\n----- LOW STOCK PRODUCTS -----")

	query := `
		SELECT id, name, price, stock, low_stock_threshold
		FROM products
		WHERE stock <= low_stock_threshold
		ORDER BY stock
	`

	rows, err := r.db.Query(context.Background(), query)
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
			&product.LowStockThreshold,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
