package repository
import("context";"q09-ecommerce-order/model";"github.com/jackc/pgx/v5/pgxpool")
type OrderRepository interface{GetOrderDetails(int)([]model.OrderDetail,error)}
type OrderRepositoryImpl struct{db *pgxpool.Pool}
func NewOrderRepository(db *pgxpool.Pool)OrderRepository{return &OrderRepositoryImpl{db}}
func(r *OrderRepositoryImpl)GetOrderDetails(id int)([]model.OrderDetail,error){rows,e:=r.db.Query(context.Background(),`SELECT o.id,c.name,p.name,oi.quantity,oi.price FROM orders o JOIN customers c ON o.customer_id=c.id JOIN order_items oi ON o.id=oi.order_id JOIN products p ON oi.product_id=p.id WHERE o.id=$1`,id);if e!=nil{return nil,e};defer rows.Close();var out []model.OrderDetail;for rows.Next(){var x model.OrderDetail;if e:=rows.Scan(&x.OrderID,&x.CustomerName,&x.ProductName,&x.Quantity,&x.Price);e!=nil{return nil,e};out=append(out,x)};return out,rows.Err()}
