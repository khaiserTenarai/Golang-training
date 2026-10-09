package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"ecommerce-order-system/model"
)

type PostgresEcommerceRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEcommerceRepository(db *pgxpool.Pool) *PostgresEcommerceRepository {
	return &PostgresEcommerceRepository{db: db}
}

func (r *PostgresEcommerceRepository) CreateCustomer(
	customer model.Customer,
) error {
	query := `
        INSERT INTO customers (name, email)
        VALUES ($1, $2)
    `

	_, err := r.db.Exec(
		context.Background(),
		query,
		customer.Name,
		customer.Email,
	)

	return err
}

func (r *PostgresEcommerceRepository) CreateProduct(
	product model.Product,
) error {
	query := `
        INSERT INTO products (name, price, stock)
        VALUES ($1, $2, $3)
    `

	_, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Stock,
	)

	return err
}

func (r *PostgresEcommerceRepository) CreateOrder(
	order model.Order,
) (int, error) {
	query := `
        INSERT INTO orders (customer_id, status)
        VALUES ($1, $2)
        RETURNING id
    `

	var orderID int

	err := r.db.QueryRow(
		context.Background(),
		query,
		order.CustomerID,
		order.Status,
	).Scan(&orderID)

	return orderID, err
}

func (r *PostgresEcommerceRepository) AddOrderItem(
	item model.OrderItem,
) error {
	query := `
        INSERT INTO order_items
        (order_id, product_id, quantity, price)
        VALUES ($1, $2, $3, $4)
    `

	_, err := r.db.Exec(
		context.Background(),
		query,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.Price,
	)

	return err
}

func (r *PostgresEcommerceRepository) GetCustomers() ([]model.Customer, error) {
	query := `
        SELECT id, name, email
        FROM customers
        ORDER BY id
    `

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []model.Customer

	for rows.Next() {
		var customer model.Customer

		if err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
		); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, rows.Err()
}

func (r *PostgresEcommerceRepository) GetProducts() ([]model.Product, error) {
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

func (r *PostgresEcommerceRepository) GetOrderDetails(
	orderID int,
) ([]model.OrderDetails, error) {
	query := `
        SELECT
            o.id,
            c.name,
            c.email,
            o.order_date,
            o.status,
            p.name,
            oi.quantity,
            oi.price,
            (oi.quantity * oi.price) AS item_total
        FROM orders o
        JOIN customers c
            ON o.customer_id = c.id
        JOIN order_items oi
            ON o.id = oi.order_id
        JOIN products p
            ON oi.product_id = p.id
        WHERE o.id = $1
        ORDER BY oi.id
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

	var details []model.OrderDetails

	for rows.Next() {
		var detail model.OrderDetails

		if err := rows.Scan(
			&detail.OrderID,
			&detail.CustomerName,
			&detail.CustomerEmail,
			&detail.OrderDate,
			&detail.Status,
			&detail.ProductName,
			&detail.Quantity,
			&detail.Price,
			&detail.ItemTotal,
		); err != nil {
			return nil, err
		}

		details = append(details, detail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(details) == 0 {
		return nil, errors.New("order not found or order has no items")
	}

	return details, nil
}

func (r *PostgresEcommerceRepository) GetAllOrderDetails() ([]model.OrderDetails, error) {
	query := `
        SELECT
            o.id,
            c.name,
            c.email,
            o.order_date,
            o.status,
            p.name,
            oi.quantity,
            oi.price,
            (oi.quantity * oi.price) AS item_total
        FROM orders o
        JOIN customers c
            ON o.customer_id = c.id
        JOIN order_items oi
            ON o.id = oi.order_id
        JOIN products p
            ON oi.product_id = p.id
        ORDER BY o.id, oi.id
    `

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []model.OrderDetails

	for rows.Next() {
		var detail model.OrderDetails

		if err := rows.Scan(
			&detail.OrderID,
			&detail.CustomerName,
			&detail.CustomerEmail,
			&detail.OrderDate,
			&detail.Status,
			&detail.ProductName,
			&detail.Quantity,
			&detail.Price,
			&detail.ItemTotal,
		); err != nil {
			return nil, err
		}

		details = append(details, detail)
	}

	return details, rows.Err()
}

// Keep fmt imported available for future repository validations.
var _ = fmt.Sprintf
