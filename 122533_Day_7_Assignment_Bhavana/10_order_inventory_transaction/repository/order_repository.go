package repository

import (
	"context"
	"example.com/q10-order-inventory/model"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	AddProduct(context.Context, string, float64, int) error
	CreateOrder(context.Context, int64, int) error
	ListProducts(context.Context) error
}
type PostgresOrderRepository struct{ db *pgxpool.Pool }

func NewPostgresOrderRepository(db *pgxpool.Pool) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: db}
}
func (r *PostgresOrderRepository) AddProduct(c context.Context, n string, p float64, s int) error {
	_, e := r.db.Exec(c, `INSERT INTO products(name,price,stock) VALUES($1,$2,$3)`, n, p, s)
	return e
}
func (r *PostgresOrderRepository) CreateOrder(c context.Context, pid int64, q int) error {
	if q <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}
	tx, e := r.db.Begin(c)
	if e != nil {
		return e
	}
	defer tx.Rollback(c)
	var stock int
	if e = tx.QueryRow(c, `SELECT stock FROM products WHERE id=$1 FOR UPDATE`, pid).Scan(&stock); e != nil {
		return fmt.Errorf("product not found")
	}
	if stock < q {
		return fmt.Errorf("insufficient stock")
	}
	var oid int64
	if e = tx.QueryRow(c, `INSERT INTO orders DEFAULT VALUES RETURNING id`).Scan(&oid); e != nil {
		return e
	}
	var price float64
	if e = tx.QueryRow(c, `SELECT price FROM products WHERE id=$1`, pid).Scan(&price); e != nil {
		return e
	}
	if _, e = tx.Exec(c, `INSERT INTO order_items(order_id,product_id,quantity,price) VALUES($1,$2,$3,$4)`, oid, pid, q, price); e != nil {
		return e
	}
	if _, e = tx.Exec(c, `UPDATE products SET stock=stock-$1 WHERE id=$2`, q, pid); e != nil {
		return e
	}
	return tx.Commit(c)
}
func (r *PostgresOrderRepository) ListProducts(c context.Context) error {
	rows, e := r.db.Query(c, `SELECT id,name,price,stock FROM products ORDER BY id`)
	if e != nil {
		return e
	}
	defer rows.Close()
	for rows.Next() {
		var p model.Product
		if e = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); e != nil {
			return e
		}
		fmt.Printf("%d | %s | %.2f | stock %d\n", p.ID, p.Name, p.Price, p.Stock)
	}
	return rows.Err()
}
