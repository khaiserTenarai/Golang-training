package repository
import("context";"errors";"q10-order-inventory-transaction/model";"github.com/jackc/pgx/v5";"github.com/jackc/pgx/v5/pgxpool")
var ErrProductNotFound=errors.New("product not found");var ErrInsufficientStock=errors.New("insufficient stock")
type OrderRepository interface{CreateOrder(model.Order)error}
type OrderRepositoryImpl struct{db *pgxpool.Pool}
func NewOrderRepository(db *pgxpool.Pool)OrderRepository{return &OrderRepositoryImpl{db}}
func(r *OrderRepositoryImpl)CreateOrder(o model.Order)error{tx,e:=r.db.Begin(context.Background());if e!=nil{return e};defer tx.Rollback(context.Background());var stock int;var price float64;e=tx.QueryRow(context.Background(),`SELECT stock,price FROM products WHERE id=$1 FOR UPDATE`,o.ProductID).Scan(&stock,&price);if e==pgx.ErrNoRows{return ErrProductNotFound};if e!=nil{return e};if stock<o.Quantity{return ErrInsufficientStock};total:=price*float64(o.Quantity);if _,e=tx.Exec(context.Background(),`UPDATE products SET stock=stock-$1 WHERE id=$2`,o.Quantity,o.ProductID);e!=nil{return e};if _,e=tx.Exec(context.Background(),`INSERT INTO orders(product_id,quantity,total)VALUES($1,$2,$3)`,o.ProductID,o.Quantity,total);e!=nil{return e};return tx.Commit(context.Background())}
