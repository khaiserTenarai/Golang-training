package repository

import (
	"context"
	"example.com/q9-ecommerce-order/model"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	AddCustomer(context.Context, string, string) error
	AddProduct(context.Context, string, float64) error
	CreateOrder(context.Context, int64, int64, int) error
	ShowOrders(context.Context) error
}
type PostgresOrderRepository struct{ db *pgxpool.Pool }

func NewPostgresOrderRepository(db *pgxpool.Pool) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: db}
}
func (r *PostgresOrderRepository) AddCustomer(ctx context.Context, n, e string) error {
	_, err := r.db.Exec(ctx, `INSERT INTO customers(name,email) VALUES($1,$2)`, n, e)
	return err
}
func (r *PostgresOrderRepository) AddProduct(ctx context.Context, n string, p float64) error {
	_, err := r.db.Exec(ctx, `INSERT INTO products(name,price) VALUES($1,$2)`, n, p)
	return err
}
func (r *PostgresOrderRepository) CreateOrder(ctx context.Context, cid, pid int64, q int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	var oid int64
	if err = tx.QueryRow(ctx, `INSERT INTO orders(customer_id) VALUES($1) RETURNING id`, cid).Scan(&oid); err != nil {
		return err
	}
	var price float64
	if err = tx.QueryRow(ctx, `SELECT price FROM products WHERE id=$1`, pid).Scan(&price); err != nil {
		return err
	}
	if q <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}
	if _, err = tx.Exec(ctx, `INSERT INTO order_items(order_id,product_id,quantity,price) VALUES($1,$2,$3,$4)`, oid, pid, q, price); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
func (r *PostgresOrderRepository) ShowOrders(ctx context.Context) error {
	rows, err := r.db.Query(ctx, `SELECT o.id,c.name,p.name,oi.quantity,oi.price,oi.quantity*oi.price FROM orders o JOIN customers c ON c.id=o.customer_id JOIN order_items oi ON oi.order_id=o.id JOIN products p ON p.id=oi.product_id ORDER BY o.id`)
	if err != nil {
		return err
	}
	defer rows.Close()
	fmt.Printf("%-8s %-18s %-18s %-8s %-10s %-10s\n", "ORDER", "CUSTOMER", "PRODUCT", "QTY", "PRICE", "TOTAL")
	for rows.Next() {
		var x model.OrderDetails
		if err := rows.Scan(&x.OrderID, &x.CustomerName, &x.ProductName, &x.Quantity, &x.Price, &x.LineTotal); err != nil {
			return err
		}
		fmt.Printf("%-8d %-18s %-18s %-8d %-10.2f %-10.2f\n", x.OrderID, x.CustomerName, x.ProductName, x.Quantity, x.Price, x.LineTotal)
	}
	return rows.Err()
}
