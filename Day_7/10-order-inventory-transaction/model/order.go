package model

import "time"

type Order struct {
	ID           int
	CustomerName string
	OrderDate    time.Time
}

type OrderItemRequest struct {
	ProductID int
	Quantity  int
}

type OrderRequest struct {
	CustomerName string
	Items        []OrderItemRequest
}

type OrderDetail struct {
	OrderID      int
	CustomerName string
	OrderDate    time.Time

	OrderItemID int

	ProductID   int
	ProductName string
	Price       float64
	Quantity    int
	ItemTotal   float64
}
