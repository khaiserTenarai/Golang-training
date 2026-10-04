package utility

import (
    "errors"
    "strings"

    "ecommerce_order_system/model"
)

func ValidateCustomer(c model.Customer) error {
    if strings.TrimSpace(c.Name) == "" {
        return errors.New("customer name is required")
    }

    if !strings.Contains(c.Email, "@") {
        return errors.New("invalid email")
    }

    return nil
}

func ValidateProduct(p model.Product) error {
    if strings.TrimSpace(p.Name) == "" {
        return errors.New("product name is required")
    }

    if p.Price < 0 {
        return errors.New("price cannot be negative")
    }

    if p.Stock < 0 {
        return errors.New("stock cannot be negative")
    }

    return nil
}

func ValidateOrder(customerID int, items []model.OrderItem) error {
    if customerID <= 0 {
        return errors.New("invalid customer id")
    }

    if len(items) == 0 {
        return errors.New("order must contain at least one product")
    }

    for _, item := range items {
        if item.ProductID <= 0 || item.Quantity <= 0 {
            return errors.New("invalid product id or quantity")
        }
    }

    return nil
}
