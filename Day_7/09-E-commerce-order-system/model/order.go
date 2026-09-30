package model
import "time"

type Order struct {
	ID         int
	CustomerID int
	OrderDate  time.Time
}

type OrderItemRequest struct {
	ProductID int
	Quantity  int
}

type OrderRequest struct {
	CustomerID int
	Items      []OrderItemRequest
}

type OrderDetail struct {
	OrderID       int
	OrderDate     time.Time
	CustomerID    int
	CustomerName  string
	CustomerEmail string
	CustomerPhone string

	OrderItemID int

	ProductID   int
	ProductName string
	Price       float64
	Quantity    int
	ItemTotal   float64
}
