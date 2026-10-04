package model

type Order struct {
	ID           int
	CustomerName string
}

type OrderItem struct {
	OrderID   int
	ProductID int
	Quantity  int
	Price     float64
}
