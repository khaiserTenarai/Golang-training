package controller

import (
    "ecommerce_order_system/model"
    "ecommerce_order_system/service"

    "github.com/jackc/pgx/v5"
)

func AddCustomer(conn *pgx.Conn, c model.Customer) error {
    return service.AddCustomer(conn, c)
}

func GetCustomers(conn *pgx.Conn) ([]model.Customer, error) {
    return service.GetCustomers(conn)
}

func AddProduct(conn *pgx.Conn, p model.Product) error {
    return service.AddProduct(conn, p)
}

func GetProducts(conn *pgx.Conn) ([]model.Product, error) {
    return service.GetProducts(conn)
}

func GetLowStockProducts(conn *pgx.Conn) ([]model.Product, error) {
    return service.GetLowStockProducts(conn)
}

func CreateOrder(conn *pgx.Conn, customerID int, items []model.OrderItem) (int, error) {
    return service.CreateOrder(conn, customerID, items)
}

func GetOrderDetails(conn *pgx.Conn) ([]model.OrderDetails, error) {
    return service.GetOrderDetails(conn)
}
