package repository
import (
	"context"
	"errors"
	"fmt"

	"order_inventory/model"

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

	// --------------------------------
	// BEGIN TRANSACTION
	// --------------------------------

	tx, err := r.db.Begin(ctx)

	if err != nil {
		return err
	}

	// If anything fails before Commit,
	// rollback everything.
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	// --------------------------------
	// STEP 1: CREATE ORDER
	// --------------------------------

	var orderID int

	err = tx.QueryRow(
		ctx,
		`
		INSERT INTO orders
		(customer_name)
		VALUES ($1)
		RETURNING id
		`,
		order.CustomerName,
	).Scan(&orderID)

	if err != nil {
		return err
	}

	// --------------------------------
	// STEP 2: PROCESS ORDER ITEMS
	// --------------------------------

	for _, item := range order.Items {

		// Lock product row until
		// transaction completes.
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

		// --------------------------------
		// STEP 3: CHECK STOCK
		// --------------------------------

		if stock < item.Quantity {

			return fmt.Errorf(
				"insufficient stock for product %d: available=%d requested=%d",
				item.ProductID,
				stock,
				item.Quantity,
			)
		}

		// --------------------------------
		// STEP 4: INSERT ORDER ITEM
		// --------------------------------

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

		// --------------------------------
		// STEP 5: REDUCE STOCK
		// --------------------------------

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

	// --------------------------------
	// STEP 6: COMMIT
	// --------------------------------

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
			o.customer_name,
			o.order_date,

			oi.id,

			p.id,
			p.name,
			p.price,

			oi.quantity,

			(p.price * oi.quantity) AS item_total

		FROM orders o

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

	details := make([]model.OrderDetail, 0)

	for rows.Next() {

		var detail model.OrderDetail

		err := rows.Scan(
			&detail.OrderID,
			&detail.CustomerName,
			&detail.OrderDate,

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

		details = append(
			details,
			detail,
		)
	}

	if len(details) == 0 {
		return nil, errors.New("order not found")
	}

	return details, rows.Err()
}

func (r *OrderRepositoryImpl) FindAll() (
	[]model.OrderDetail,
	error,
) {

	query := `
		SELECT
			o.id,
			o.customer_name,
			o.order_date,

			oi.id,

			p.id,
			p.name,
			p.price,

			oi.quantity,

			(p.price * oi.quantity) AS item_total

		FROM orders o

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

	details := make([]model.OrderDetail, 0)

	for rows.Next() {

		var detail model.OrderDetail

		err := rows.Scan(
			&detail.OrderID,
			&detail.CustomerName,
			&detail.OrderDate,

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

		details = append(
			details,
			detail,
		)
	}

	return details, rows.Err()
}
