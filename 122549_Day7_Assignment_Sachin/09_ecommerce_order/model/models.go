package model

type Customer struct {
	ID    int64
	Name  string
	Email string
}
type Product struct {
	ID    int64
	Name  string
	Price float64
}
type OrderItem struct {
	ProductID   int64
	ProductName string
	Quantity    int
	Price       float64
}
type OrderDetails struct {
	OrderID      int64
	CustomerName string
	ProductName  string
	Quantity     int
	Price        float64
	LineTotal    float64
}
