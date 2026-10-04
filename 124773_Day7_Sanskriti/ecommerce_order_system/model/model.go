package model

type Customer struct {
    ID    int
    Name  string
    Email string
}

type Product struct {
    ID            int
    Name          string
    Price         float64
    Stock         int
    LowStockLimit int
}

type OrderItem struct {
    ProductID int
    Quantity  int
}

type Order struct {
    ID         int
    CustomerID int
    Total      float64
}

type OrderDetails struct {
    OrderID      int
    CustomerName string
    ProductName  string
    Quantity     int
    UnitPrice    float64
    LineTotal    float64
    OrderTotal   float64
}
