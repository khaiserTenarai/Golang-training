package repository

import (
	"context"
	"errors"

	"product_inventory/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

func (r *ProductRepositoryImpl) Save(
	product model.Product,
) error {

	query := `
		INSERT INTO products
		(name, price, quantity)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Quantity,
	)

	return err
}

func (r *ProductRepositoryImpl) FindByID(
	id int,
) (model.Product, error) {

	var product model.Product

	query := `
		SELECT id, name, price, quantity
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
		&product.Quantity,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return product, errors.New("product not found")
	}

	return product, err
}

func (r *ProductRepositoryImpl) FindAll() (
	[]model.Product,
	error,
) {

	query := `
		SELECT id, name, price, quantity
		FROM products
		ORDER BY id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {

		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Quantity,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, rows.Err()
}

func (r *ProductRepositoryImpl) Update(
	product model.Product,
) error {

	query := `
		UPDATE products
		SET name = $1,
		    price = $2,
		    quantity = $3
		WHERE id = $4
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Quantity,
		product.ID,
	)

	return err
}

func (r *ProductRepositoryImpl) Delete(
	id int,
) error {

	query := `
		DELETE FROM products
		WHERE id = $1
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}

func (r *ProductRepositoryImpl) IncreaseStock(
	id int,
	quantity int,
) error {

	query := `
		UPDATE products
		SET quantity = quantity + $1
		WHERE id = $2
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		quantity,
		id,
	)

	return err
}

func (r *ProductRepositoryImpl) DecreaseStock(
	id int,
	quantity int,
) error {

	query := `
		UPDATE products
		SET quantity = quantity - $1
		WHERE id = $2
		  AND quantity >= $1
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
		return errors.New("not enough stock or product not found")
	}

	return nil
}

func (r *ProductRepositoryImpl) FindLowStock(
	limit int,
) ([]model.Product, error) {

	query := `
		SELECT id, name, price, quantity
		FROM products
		WHERE quantity <= $1
		ORDER BY quantity
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		limit,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {

		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Quantity,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, rows.Err()
}
