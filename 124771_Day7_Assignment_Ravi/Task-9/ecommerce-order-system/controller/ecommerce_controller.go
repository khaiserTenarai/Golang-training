package controller

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "ecommerce-order-system/model"
    "eecommerce-order-system/service"
)

type EcommerceController struct {
    service service.EcommerceService
    reader  *bufio.Reader
}

func NewEcommerceController(
    service service.EcommerceService,
) *EcommerceController {
    return &EcommerceController{
        service: service,
        reader:  bufio.NewReader(os.Stdin),
    }
}

func (c *EcommerceController) Start() {
    for {
        fmt.Println()
        fmt.Println("========================================")
        fmt.Println("       E-COMMERCE ORDER SYSTEM")
        fmt.Println("========================================")
        fmt.Println("1. Create Customer")
        fmt.Println("2. Create Product")
        fmt.Println("3. Create Order")
        fmt.Println("4. View Customers")
        fmt.Println("5. View Products")
        fmt.Println("6. View Order Details")
        fmt.Println("7. View All Order Details")
        fmt.Println("8. Exit")
        fmt.Println("========================================")

        choice := c.readInt("Enter choice: ")

        switch choice {
        case 1:
            c.createCustomer()
        case 2:
            c.createProduct()
        case 3:
            c.createOrder()
        case 4:
            c.viewCustomers()
        case 5:
            c.viewProducts()
        case 6:
            c.viewOrderDetails()
        case 7:
            c.viewAllOrderDetails()
        case 8:
            fmt.Println("Thank you for using E-Commerce Order System.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
}

func (c *EcommerceController) createCustomer() {
    customer := model.Customer{
        Name:  c.readString("Enter customer name: "),
        Email: c.readString("Enter customer email: "),
    }

    if err := c.service.CreateCustomer(customer); err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Customer created successfully.")
}

func (c *EcommerceController) createProduct() {
    product := model.Product{
        Name:  c.readString("Enter product name: "),
        Price: c.readFloat("Enter product price: "),
        Stock: c.readInt("Enter product stock: "),
    }

    if err := c.service.CreateProduct(product); err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Product created successfully.")
}

func (c *EcommerceController) createOrder() {
    customerID := c.readInt("Enter customer ID: ")
    itemCount := c.readInt("Enter number of products in order: ")

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
        price := c.readFloat("Enter product price: ")

        items = append(items, model.OrderItem{
            ProductID: productID,
            Quantity:  quantity,
            Price:     price,
        })
    }

    order := model.Order{
        CustomerID: customerID,
        Status:     "PLACED",
    }

    orderID, err := c.service.CreateOrder(order, items)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Order created successfully.")
    fmt.Println("Order ID:", orderID)
}

func (c *EcommerceController) viewCustomers() {
    customers, err := c.service.GetCustomers()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println()
    fmt.Println("ID | Name | Email")
    fmt.Println("-----------------------------------------")

    for _, customer := range customers {
        fmt.Printf(
            "%d | %s | %s
",
            customer.ID,
            customer.Name,
            customer.Email,
        )
    }
}

func (c *EcommerceController) viewProducts() {
    products, err := c.service.GetProducts()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println()
    fmt.Println("ID | Product | Price | Stock")
    fmt.Println("-----------------------------------------")

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

func (c *EcommerceController) viewOrderDetails() {
    orderID := c.readInt("Enter order ID: ")

    details, err := c.service.GetOrderDetails(orderID)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    printOrderDetails(details)
}

func (c *EcommerceController) viewAllOrderDetails() {
    details, err := c.service.GetAllOrderDetails()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(details) == 0 {
        fmt.Println("No orders found.")
        return
    }

    currentOrderID := -1

    for _, detail := range details {
        if detail.OrderID != currentOrderID {
            if currentOrderID != -1 {
                fmt.Println("----------------------------------------")
            }

            currentOrderID = detail.OrderID

            fmt.Printf("
Order ID: %d
", detail.OrderID)
            fmt.Println("Customer:", detail.CustomerName)
            fmt.Println("Email:", detail.CustomerEmail)
            fmt.Println("Date:", detail.OrderDate.Format("2006-01-02 15:04:05"))
            fmt.Println("Status:", detail.Status)
            fmt.Println("Items:")
        }

        fmt.Printf(
            "  %s | Qty: %d | Price: ₹%.2f | Total: ₹%.2f
",
            detail.ProductName,
            detail.Quantity,
            detail.Price,
            detail.ItemTotal,
        )
    }
}

func printOrderDetails(details []model.OrderDetails) {
    if len(details) == 0 {
        fmt.Println("No order details found.")
        return
    }

    first := details[0]

    fmt.Println()
    fmt.Println("========================================")
    fmt.Println("           COMPLETE ORDER")
    fmt.Println("========================================")
    fmt.Println("Order ID:", first.OrderID)
    fmt.Println("Customer:", first.CustomerName)
    fmt.Println("Email:", first.CustomerEmail)
    fmt.Println("Date:", first.OrderDate.Format("2006-01-02 15:04:05"))
    fmt.Println("Status:", first.Status)
    fmt.Println()
    fmt.Println("Products:")

    total := 0.0

    for _, detail := range details {
        fmt.Printf(
            "%s | Qty: %d | Price: ₹%.2f | Total: ₹%.2f
",
            detail.ProductName,
            detail.Quantity,
            detail.Price,
            detail.ItemTotal,
        )

        total += detail.ItemTotal
    }

    fmt.Println("----------------------------------------")
    fmt.Printf("Order Total: ₹%.2f
", total)
}

func (c *EcommerceController) readString(prompt string) string {
    fmt.Print(prompt)
    value, _ := c.reader.ReadString('
')
    return strings.TrimSpace(value)
}

func (c *EcommerceController) readInt(prompt string) int {
    for {
        value := c.readString(prompt)

        number, err := strconv.Atoi(value)
        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid number.")
    }
}

func (c *EcommerceController) readFloat(prompt string) float64 {
    for {
        value := c.readString(prompt)

        number, err := strconv.ParseFloat(value, 64)
        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid amount.")
    }
}
