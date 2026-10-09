package model

type OrderRequest struct {
    CustomerName string
    Items        []OrderItem
}
