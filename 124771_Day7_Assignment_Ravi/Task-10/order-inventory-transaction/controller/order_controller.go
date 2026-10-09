package controller

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "order-inventory-transaction/model"
    "order-inventory-transaction/service"
)

type OrderController struct {
    service service.OrderService
    reader  *bufio.Reader
}

func NewOrderController(
    service service.OrderService,
) *OrderController {
    return &OrderController{
        service: service,
        reader:  bufio.NewReader(os.Stdin),
    }
}

func (c *OrderController) Start() {
    for {
        fmt.Println()
        fmt.Println("========================================")
        fmt.Println("     ORDER & INVENTORY TRANSACTION")
        fmt.Println("========================================")
        fmt.Println("1. View Products")
        fmt.Println("2. Create Order")
        fmt.Println("3. Exit")
        fmt.Println("========================================")

        choice := c.readInt("Enter choice: ")

        switch choice {
        case 1:
            c.viewProducts()
        case 2:
            c.createOrder()
        case 3:
            fmt.Println("Thank you for using the system.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
}

func (c *OrderController) viewProducts() {
    products, err := c.service.GetProducts()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println()
    fmt.Println("ID | Product | Price | Stock")
    fmt.Println("-------------------------------------------")

    for _, product := range products {
        fmt.Printf(
            "%d | %s | ₹%.2f | %d
",
            product.ID,
            product.Name,
            product.Price,
            product.Stock,
        )
    }
}

func (c *OrderController) createOrder() {
    customerName := c.readString("Enter customer name: ")
    itemCount := c.readInt("Enter number of different products: ")

    if itemCount <= 0 {
        fmt.Println("Order must contain at least one product.")
        return
    }

    items := make([]model.OrderItem, 0, itemCount)

    for i := 0; i < itemCount; i++ {
        fmt.Printf("
Product %d
", i+1)

        productID := c.readInt("Enter product ID: ")
        quantity := c.readInt("Enter quantity: ")

        items = append(items, model.OrderItem{
            ProductID: productID,
            Quantity:  quantity,
        })
    }

    request := model.OrderRequest{
        CustomerName: customerName,
        Items:        items,
    }

    fmt.Println()
    fmt.Println("Creating order and updating inventory...")

    orderID, err := c.service.CreateOrder(request)
    if err != nil {
        fmt.Println("Order failed:", err)
        fmt.Println("Complete operation was rolled back.")
        return
    }

    fmt.Println("Order created successfully.")
    fmt.Println("Order ID:", orderID)
    fmt.Println("Inventory updated successfully.")
    fmt.Println("Transaction committed.")
}

func (c *OrderController) readString(prompt string) string {
    fmt.Print(prompt)
    value, _ := c.reader.ReadString('
')
    return strings.TrimSpace(value)
}

func (c *OrderController) readInt(prompt string) int {
    for {
        value := c.readString(prompt)

        number, err := strconv.Atoi(value)
        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid number.")
    }
}
