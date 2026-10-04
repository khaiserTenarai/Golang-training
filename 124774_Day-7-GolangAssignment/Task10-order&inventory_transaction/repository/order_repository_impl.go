package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewOrderRepository(
	db *pgxpool.Pool,
) OrderRepository {

	return &OrderRepositoryImpl{
		db: db,
	}
}

func (r *OrderRepositoryImpl) CreateOrder(
	customerName string,
	productID int,
	quantity int,
) error {

	ctx := context.Background()

	// Start transaction
	tx, err := r.db.Begin(ctx)

	if err != nil {
		return err
	}

	// Rollback if anything fails
	defer tx.Rollback(ctx)

	// Get product price and stock
	var price float64
	var stock int

	err = tx.QueryRow(
		ctx,
		`SELECT price, stock
		 FROM products
		 WHERE id = $1
		 FOR UPDATE`,
		productID,
	).Scan(&price, &stock)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("product not found")
		}

		return err
	}

	// Check stock
	if stock < quantity {
		return errors.New("insufficient stock")
	}

	// Create order
	var orderID int

	err = tx.QueryRow(
		ctx,
		`INSERT INTO orders (customer_name)
		 VALUES ($1)
		 RETURNING id`,
		customerName,
	).Scan(&orderID)

	if err != nil {
		return err
	}

	// Create order item
	_, err = tx.Exec(
		ctx,
		`INSERT INTO order_items
		 (order_id, product_id, quantity, price)
		 VALUES ($1, $2, $3, $4)`,
		orderID,
		productID,
		quantity,
		price,
	)

	if err != nil {
		return err
	}

	// Reduce stock
	_, err = tx.Exec(
		ctx,
		`UPDATE products
		 SET stock = stock - $1
		 WHERE id = $2`,
		quantity,
		productID,
	)

	if err != nil {
		return err
	}

	// Commit
	return tx.Commit(ctx)
}
