package repository

import (
	"context"
	"errors"

	"ecommerce-order/model"

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

// Create Customer

func (r *OrderRepositoryImpl) CreateCustomer(
	customer model.Customer,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO customers (name, email)
		 VALUES ($1, $2)`,
		customer.Name,
		customer.Email,
	)

	return err
}

// Create Product

func (r *OrderRepositoryImpl) CreateProduct(
	product model.Product,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO products (name, price)
		 VALUES ($1, $2)`,
		product.Name,
		product.Price,
	)

	return err
}

// Create Order

func (r *OrderRepositoryImpl) CreateOrder(
	customerID int,
) (int, error) {

	var orderID int

	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO orders (customer_id)
		 VALUES ($1)
		 RETURNING id`,
		customerID,
	).Scan(&orderID)

	return orderID, err
}

// Create Order Item

func (r *OrderRepositoryImpl) CreateOrderItem(
	item model.OrderItem,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO order_items
		 (order_id, product_id, quantity, price)
		 VALUES ($1, $2, $3, $4)`,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.Price,
	)

	return err
}

// JOIN query

func (r *OrderRepositoryImpl) GetOrderDetails(
	orderID int,
) ([]model.OrderDetail, error) {

	query := `
		SELECT
			o.id,
			c.name,
			p.name,
			oi.quantity,
			oi.price
		FROM orders o
		JOIN customers c
			ON o.customer_id = c.id
		JOIN order_items oi
			ON o.id = oi.order_id
		JOIN products p
			ON oi.product_id = p.id
		WHERE o.id = $1
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		orderID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var details []model.OrderDetail

	for rows.Next() {

		var detail model.OrderDetail

		err := rows.Scan(
			&detail.OrderID,
			&detail.CustomerName,
			&detail.ProductName,
			&detail.Quantity,
			&detail.Price,
		)

		if err != nil {
			return nil, err
		}

		details = append(details, detail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(details) == 0 {
		return nil, errors.New("order not found")
	}

	return details, nil
}
