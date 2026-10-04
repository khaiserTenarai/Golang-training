package repository

import (
    "context"
    "fmt"

    "ecommerce_order_system/model"
    "github.com/jackc/pgx/v5"
)

func AddCustomer(conn *pgx.Conn, c model.Customer) error {
    _, err := conn.Exec(context.Background(),
        "INSERT INTO customers(name, email) VALUES($1, $2)",
        c.Name, c.Email,
    )
    return err
}

func GetCustomers(conn *pgx.Conn) ([]model.Customer, error) {
    rows, err := conn.Query(context.Background(),
        "SELECT id, name, email FROM customers ORDER BY id")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var customers []model.Customer

    for rows.Next() {
        var c model.Customer

        err := rows.Scan(&c.ID, &c.Name, &c.Email)
        if err != nil {
            return nil, err
        }

        customers = append(customers, c)
    }

    return customers, rows.Err()
}

func AddProduct(conn *pgx.Conn, p model.Product) error {
    _, err := conn.Exec(context.Background(),
        `INSERT INTO products(name, price, stock, low_stock_limit)
         VALUES($1, $2, $3, $4)`,
        p.Name, p.Price, p.Stock, p.LowStockLimit,
    )
    return err
}

func GetProducts(conn *pgx.Conn) ([]model.Product, error) {
    rows, err := conn.Query(context.Background(),
        `SELECT id, name, price, stock, low_stock_limit
         FROM products ORDER BY id`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []model.Product

    for rows.Next() {
        var p model.Product

        err := rows.Scan(
            &p.ID,
            &p.Name,
            &p.Price,
            &p.Stock,
            &p.LowStockLimit,
        )
        if err != nil {
            return nil, err
        }

        products = append(products, p)
    }

    return products, rows.Err()
}

func GetLowStockProducts(conn *pgx.Conn) ([]model.Product, error) {
    rows, err := conn.Query(context.Background(),
        `SELECT id, name, price, stock, low_stock_limit
         FROM products
         WHERE stock <= low_stock_limit
         ORDER BY stock`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []model.Product

    for rows.Next() {
        var p model.Product

        err := rows.Scan(
            &p.ID,
            &p.Name,
            &p.Price,
            &p.Stock,
            &p.LowStockLimit,
        )
        if err != nil {
            return nil, err
        }

        products = append(products, p)
    }

    return products, rows.Err()
}

// CreateOrder uses one transaction for the complete operation.
func CreateOrder(conn *pgx.Conn, customerID int, items []model.OrderItem) (int, error) {

    tx, err := conn.Begin(context.Background())
    if err != nil {
        return 0, err
    }

    // If anything fails, undo everything.
    defer tx.Rollback(context.Background())

    var total float64

    // Check stock before creating the order.
    for _, item := range items {

        var price float64
        var stock int

        err := tx.QueryRow(
            context.Background(),
            `SELECT price, stock
             FROM products
             WHERE id = $1
             FOR UPDATE`,
            item.ProductID,
        ).Scan(&price, &stock)

        if err != nil {
            return 0, err
        }

        if stock < item.Quantity {
            return 0, fmt.Errorf(
                "not enough stock for product %d. Available: %d",
                item.ProductID,
                stock,
            )
        }

        total += price * float64(item.Quantity)
    }

    // Create order.
    var orderID int

    err = tx.QueryRow(
        context.Background(),
        `INSERT INTO orders(customer_id, total_amount)
         VALUES($1, $2)
         RETURNING id`,
        customerID,
        total,
    ).Scan(&orderID)

    if err != nil {
        return 0, err
    }

    // Add order items and reduce stock.
    for _, item := range items {

        var price float64

        err := tx.QueryRow(
            context.Background(),
            "SELECT price FROM products WHERE id = $1",
            item.ProductID,
        ).Scan(&price)

        if err != nil {
            return 0, err
        }

        _, err = tx.Exec(
            context.Background(),
            `INSERT INTO order_items(order_id, product_id, quantity, unit_price)
             VALUES($1, $2, $3, $4)`,
            orderID,
            item.ProductID,
            item.Quantity,
            price,
        )

        if err != nil {
            return 0, err
        }

        _, err = tx.Exec(
            context.Background(),
            `UPDATE products
             SET stock = stock - $1
             WHERE id = $2`,
            item.Quantity,
            item.ProductID,
        )

        if err != nil {
            return 0, err
        }
    }

    err = tx.Commit(context.Background())
    if err != nil {
        return 0, err
    }

    return orderID, nil
}

func GetOrderDetails(conn *pgx.Conn) ([]model.OrderDetails, error) {

    rows, err := conn.Query(context.Background(),
        `SELECT
            o.id,
            c.name,
            p.name,
            oi.quantity,
            oi.unit_price,
            oi.quantity * oi.unit_price,
            o.total_amount
         FROM orders o
         JOIN customers c ON c.id = o.customer_id
         JOIN order_items oi ON oi.order_id = o.id
         JOIN products p ON p.id = oi.product_id
         ORDER BY o.id`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var details []model.OrderDetails

    for rows.Next() {
        var d model.OrderDetails

        err := rows.Scan(
            &d.OrderID,
            &d.CustomerName,
            &d.ProductName,
            &d.Quantity,
            &d.UnitPrice,
            &d.LineTotal,
            &d.OrderTotal,
        )

        if err != nil {
            return nil, err
        }

        details = append(details, d)
    }

    return details, rows.Err()
}
