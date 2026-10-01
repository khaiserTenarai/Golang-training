package model

import "fmt"

type Product struct {
    ID            int64   `json:"id"`
    Name          string  `json:"name"`
    Category      string  `json:"category"`
    Price         float64 `json:"price"`
    StockQuantity int     `json:"stock_quantity"`
    ReorderLevel  int     `json:"reorder_level"`
}

func (p Product) Display() {
    fmt.Println("----------------------------------------")
    fmt.Println("ID             :", p.ID)
    fmt.Println("Name           :", p.Name)
    fmt.Println("Category       :", p.Category)
    fmt.Printf("Price          : $%.2f\n", p.Price)
    fmt.Println("Stock Quantity :", p.StockQuantity)
    fmt.Println("Reorder Level  :", p.ReorderLevel)
    fmt.Println("----------------------------------------")
}