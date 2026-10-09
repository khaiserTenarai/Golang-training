package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"product-inventory/model"
)

var ErrProductNotFound = errors.New(
	"product not found",
)

type PostgresProductRepository struct {
	db *pgxpool.Pool
}

func NewPostgresProductRepository(
	db *pgxpool.Pool,
) *PostgresProductRepository {

	return &PostgresProductRepository{
		db: db,
	}
}

// ----------------------------------------------------
// Create
// ----------------------------------------------------

func (r *PostgresProductRepository) Create(
	ctx context.Context,
	product *model.Product,
) (*model.Product, error) {

	var result model.Product

	err := r.db.QueryRow(
		ctx,
		`
		INSERT INTO public.products (
			name,
			description,
			price,
			stock_quantity,
			low_stock_threshold
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			name,
			description,
			price,
			stock_quantity,
			low_stock_threshold,
			created_at,
			updated_at
		`,
		product.Name,
		product.Description,
		product.Price,
		product.StockQuantity,
		product.LowStockThreshold,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Description,
		&result.Price,
		&result.StockQuantity,
		&result.LowStockThreshold,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to create product: %w",
			err,
		)
	}

	return &result, nil
}

// ----------------------------------------------------
// Get By ID
// ----------------------------------------------------

func (r *PostgresProductRepository) GetByID(
	ctx context.Context,
	id int64,
) (*model.Product, error) {

	var product model.Product

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			name,
			description,
			price,
			stock_quantity,
			low_stock_threshold,
			created_at,
			updated_at
		FROM public.products
		WHERE id = $1
		`,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.StockQuantity,
		&product.LowStockThreshold,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotFound
		}

		return nil, fmt.Errorf(
			"failed to get product: %w",
			err,
		)
	}

	return &product, nil
}

// ----------------------------------------------------
// Get All
// ----------------------------------------------------

func (r *PostgresProductRepository) GetAll(
	ctx context.Context,
) ([]model.Product, error) {

	rows, err := r.db.Query(
		ctx,
		`
		SELECT
			id,
			name,
			description,
			price,
			stock_quantity,
			low_stock_threshold,
			created_at,
			updated_at
		FROM public.products
		ORDER BY id
		`,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to get products: %w",
			err,
		)
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
			&product.StockQuantity,
			&product.LowStockThreshold,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"failed to scan product: %w",
				err,
			)
		}

		products = append(
			products,
			product,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"failed to read products: %w",
			err,
		)
	}

	return products, nil
}

// ----------------------------------------------------
// Update
// ----------------------------------------------------

func (r *PostgresProductRepository) Update(
	ctx context.Context,
	product *model.Product,
) (*model.Product, error) {

	var result model.Product

	err := r.db.QueryRow(
		ctx,
		`
		UPDATE public.products
		SET
			name = $1,
			description = $2,
			price = $3,
			low_stock_threshold = $4,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $5
		RETURNING
			id,
			name,
			description,
			price,
			stock_quantity,
			low_stock_threshold,
			created_at,
			updated_at
		`,
		product.Name,
		product.Description,
		product.Price,
		product.LowStockThreshold,
		product.ID,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Description,
		&result.Price,
		&result.StockQuantity,
		&result.LowStockThreshold,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotFound
		}

		return nil, fmt.Errorf(
			"failed to update product: %w",
			err,
		)
	}

	return &result, nil
}

// ----------------------------------------------------
// Delete
// ----------------------------------------------------

func (r *PostgresProductRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	result, err := r.db.Exec(
		ctx,
		`
		DELETE FROM public.products
		WHERE id = $1
		`,
		id,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to delete product: %w",
			err,
		)
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	return nil
}

// ----------------------------------------------------
// Get By ID For Update
// ----------------------------------------------------

func (r *PostgresProductRepository) GetByIDForUpdate(
	ctx context.Context,
	tx pgx.Tx,
	id int64,
) (*model.Product, error) {

	var product model.Product

	err := tx.QueryRow(
		ctx,
		`
		SELECT
			id,
			name,
			description,
			price,
			stock_quantity,
			low_stock_threshold,
			created_at,
			updated_at
		FROM public.products
		WHERE id = $1
		FOR UPDATE
		`,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.StockQuantity,
		&product.LowStockThreshold,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotFound
		}

		return nil, fmt.Errorf(
			"failed to lock product: %w",
			err,
		)
	}

	return &product, nil
}

// ----------------------------------------------------
// Update Stock
// ----------------------------------------------------

func (r *PostgresProductRepository) UpdateStock(
	ctx context.Context,
	tx pgx.Tx,
	id int64,
	stockQuantity int,
) error {

	_, err := tx.Exec(
		ctx,
		`
		UPDATE public.products
		SET
			stock_quantity = $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
		`,
		stockQuantity,
		id,
	)

	if err != nil {
		return fmt.Errorf(
			"failed to update stock: %w",
			err,
		)
	}

	return nil
}

// ----------------------------------------------------
// Low Stock Products
// ----------------------------------------------------

func (r *PostgresProductRepository) GetLowStockProducts(
	ctx context.Context,
) ([]model.Product, error) {

	rows, err := r.db.Query(
		ctx,
		`
		SELECT
			id,
			name,
			description,
			price,
			stock_quantity,
			low_stock_threshold,
			created_at,
			updated_at
		FROM public.products
		WHERE stock_quantity <= low_stock_threshold
		ORDER BY stock_quantity ASC, id ASC
		`,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to get low stock products: %w",
			err,
		)
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
			&product.StockQuantity,
			&product.LowStockThreshold,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"failed to scan low stock product: %w",
				err,
			)
		}

		products = append(
			products,
			product,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
