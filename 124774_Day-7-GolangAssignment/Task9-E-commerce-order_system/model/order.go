package model

type Order struct {
	ID         int
	CustomerID int
}

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
	Price     float64
}

type OrderDetail struct {
	OrderID      int
	CustomerName string
	ProductName  string
	Quantity     int
	Price        float64
}
