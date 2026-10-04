package service

import (
    "ecommerce_order_system/model"
    "ecommerce_order_system/repository"
    "ecommerce_order_system/utility"

    "github.com/jackc/pgx/v5"
)

func AddCustomer(conn *pgx.Conn, c model.Customer) error {
    if err := utility.ValidateCustomer(c); err != nil {
        return err
    }

    return repository.AddCustomer(conn, c)
}

func GetCustomers(conn *pgx.Conn) ([]model.Customer, error) {
    return repository.GetCustomers(conn)
}

func AddProduct(conn *pgx.Conn, p model.Product) error {
    if err := utility.ValidateProduct(p); err != nil {
        return err
    }

    return repository.AddProduct(conn, p)
}

func GetProducts(conn *pgx.Conn) ([]model.Product, error) {
    return repository.GetProducts(conn)
}

func GetLowStockProducts(conn *pgx.Conn) ([]model.Product, error) {
    return repository.GetLowStockProducts(conn)
}

func CreateOrder(conn *pgx.Conn, customerID int, items []model.OrderItem) (int, error) {
    if err := utility.ValidateOrder(customerID, items); err != nil {
        return 0, err
    }

    return repository.CreateOrder(conn, customerID, items)
}

func GetOrderDetails(conn *pgx.Conn) ([]model.OrderDetails, error) {
    return repository.GetOrderDetails(conn)
}
