package repository
import (
	"context"
	"errors"
	"fmt"

	"ecommerce/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	CreateOrder(order model.OrderRequest) error
	FindByID(id int) ([]model.OrderDetail, error)
	FindAll() ([]model.OrderDetail, error)
}

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
	order model.OrderRequest,
) error {

	ctx := context.Background()

	tx, err := r.db.Begin(ctx)

	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	// Check customer
	var customerID int

	err = tx.QueryRow(
		ctx,
		`SELECT id FROM customers WHERE id = $1`,
		order.CustomerID,
	).Scan(&customerID)

	if errors.Is(err, pgx.ErrNoRows) {
		return errors.New("customer not found")
	}

	if err != nil {
		return err
	}

	// Create order
	var orderID int

	err = tx.QueryRow(
		ctx,
		`
		INSERT INTO orders (customer_id)
		VALUES ($1)
		RETURNING id
		`,
		order.CustomerID,
	).Scan(&orderID)

	if err != nil {
		return err
	}

	// Create order items
	for _, item := range order.Items {

		var stock int

		err = tx.QueryRow(
			ctx,
			`
			SELECT stock
			FROM products
			WHERE id = $1
			FOR UPDATE
			`,
			item.ProductID,
		).Scan(&stock)

		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf(
				"product %d not found",
				item.ProductID,
			)
		}

		if err != nil {
			return err
		}

		if stock < item.Quantity {
			return fmt.Errorf(
				"insufficient stock for product %d",
				item.ProductID,
			)
		}

		_, err = tx.Exec(
			ctx,
			`
			INSERT INTO order_items
			(order_id, product_id, quantity)
			VALUES ($1, $2, $3)
			`,
			orderID,
			item.ProductID,
			item.Quantity,
		)

		if err != nil {
			return err
		}

		// Reduce stock
		_, err = tx.Exec(
			ctx,
			`
			UPDATE products
			SET stock = stock - $1
			WHERE id = $2
			`,
			item.Quantity,
			item.ProductID,
		)

		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepositoryImpl) FindByID(
	id int,
) ([]model.OrderDetail, error) {

	query := `
		SELECT
			o.id,
			o.order_date,

			c.id,
			c.name,
			c.email,
			c.phone,

			oi.id,

			p.id,
			p.name,
			p.price,

			oi.quantity,

			(p.price * oi.quantity) AS item_total

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
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orderDetails := make([]model.OrderDetail, 0)

	for rows.Next() {

		var detail model.OrderDetail

		err := rows.Scan(
			&detail.OrderID,
			&detail.OrderDate,

			&detail.CustomerID,
			&detail.CustomerName,
			&detail.CustomerEmail,
			&detail.CustomerPhone,

			&detail.OrderItemID,

			&detail.ProductID,
			&detail.ProductName,
			&detail.Price,

			&detail.Quantity,
			&detail.ItemTotal,
		)

		if err != nil {
			return nil, err
		}

		orderDetails = append(
			orderDetails,
			detail,
		)
	}

	if len(orderDetails) == 0 {
		return nil, errors.New("order not found")
	}

	return orderDetails, rows.Err()
}

func (r *OrderRepositoryImpl) FindAll() (
	[]model.OrderDetail,
	error,
) {

	query := `
		SELECT
			o.id,
			o.order_date,

			c.id,
			c.name,
			c.email,
			c.phone,

			oi.id,

			p.id,
			p.name,
			p.price,

			oi.quantity,

			(p.price * oi.quantity) AS item_total

		FROM orders o

		JOIN customers c
			ON o.customer_id = c.id

		JOIN order_items oi
			ON o.id = oi.order_id

		JOIN products p
			ON oi.product_id = p.id

		ORDER BY o.id, oi.id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orderDetails := make([]model.OrderDetail, 0)

	for rows.Next() {

		var detail model.OrderDetail

		err := rows.Scan(
			&detail.OrderID,
			&detail.OrderDate,

			&detail.CustomerID,
			&detail.CustomerName,
			&detail.CustomerEmail,
			&detail.CustomerPhone,

			&detail.OrderItemID,

			&detail.ProductID,
			&detail.ProductName,
			&detail.Price,

			&detail.Quantity,
			&detail.ItemTotal,
		)

		if err != nil {
			return nil, err
		}

		orderDetails = append(
			orderDetails,
			detail,
		)
	}

	return orderDetails, rows.Err()
}
