package repository

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"order-inventory-transaction/model"
)

type PostgresOrderRepository struct {
	db *pgxpool.Pool
}

func NewPostgresOrderRepository(db *pgxpool.Pool) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) GetProducts() ([]model.Product, error) {
	query := `
        SELECT id, name, price, stock
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

		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, rows.Err()
}

func (r *PostgresOrderRepository) GetProduct(
	productID int,
) (model.Product, error) {
	query := `
        SELECT id, name, price, stock
        FROM products
        WHERE id = $1
    `

	var product model.Product

	err := r.db.QueryRow(
		context.Background(),
		query,
		productID,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.Product{}, fmt.Errorf(
			"product with ID %d not found",
			productID,
		)
	}

	return product, err
}

func (r *PostgresOrderRepository) CreateOrder(
	order model.Order,
	items []model.OrderItem,
) (int, error) {
	ctx := context.Background()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("unable to begin transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	quantities := make(map[int]int)
	for _, item := range items {
		quantities[item.ProductID] += item.Quantity
	}

	productIDs := make([]int, 0, len(quantities))
	for productID := range quantities {
		productIDs = append(productIDs, productID)
	}

	// Lock products in a consistent ID order to reduce deadlock risk.
	sort.Ints(productIDs)

	lockedProducts := make(map[int]model.Product)

	for _, productID := range productIDs {
		query := `
            SELECT id, name, price, stock
            FROM products
            WHERE id = $1
            FOR UPDATE
        `

		var product model.Product

		err := tx.QueryRow(
			ctx,
			query,
			productID,
		).Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
		)

		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf(
				"product with ID %d not found",
				productID,
			)
		}

		if err != nil {
			return 0, fmt.Errorf(
				"unable to lock product %d: %w",
				productID,
				err,
			)
		}

		lockedProducts[productID] = product
	}

	// Check stock BEFORE creating the order.
	for productID, quantity := range quantities {
		product := lockedProducts[productID]

		if product.Stock < quantity {
			return 0, fmt.Errorf(
				"insufficient stock for %s: available %d, requested %d",
				product.Name,
				product.Stock,
				quantity,
			)
		}
	}

	// Create the order only after all stock checks succeed.
	orderQuery := `
        INSERT INTO orders (customer_name, status)
        VALUES ($1, $2)
        RETURNING id
    `

	var orderID int

	err = tx.QueryRow(
		ctx,
		orderQuery,
		order.CustomerName,
		order.Status,
	).Scan(&orderID)

	if err != nil {
		return 0, fmt.Errorf(
			"unable to create order: %w",
			err,
		)
	}

	// Insert order items and reduce stock.
	for _, item := range items {
		product := lockedProducts[item.ProductID]

		itemQuery := `
            INSERT INTO order_items
            (order_id, product_id, quantity, price)
            VALUES ($1, $2, $3, $4)
        `

		if _, err := tx.Exec(
			ctx,
			itemQuery,
			orderID,
			item.ProductID,
			item.Quantity,
			product.Price,
		); err != nil {
			return 0, fmt.Errorf(
				"unable to create order item: %w",
				err,
			)
		}

		stockQuery := `
            UPDATE products
            SET stock = stock - $1
            WHERE id = $2
        `

		if _, err := tx.Exec(
			ctx,
			stockQuery,
			item.Quantity,
			item.ProductID,
		); err != nil {
			return 0, fmt.Errorf(
				"unable to reduce stock for product %d: %w",
				item.ProductID,
				err,
			)
		}
	}

	// All operations succeeded.
	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf(
			"unable to commit order transaction: %w",
			err,
		)
	}

	return orderID, nil
}
