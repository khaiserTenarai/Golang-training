package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ecommerce_order_system/controller"
	"ecommerce_order_system/model"

	"github.com/jackc/pgx/v5"
)

func Start(conn *pgx.Conn) {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println()
		fmt.Println("===== E-COMMERCE MANAGEMENT =====")
		fmt.Println("1. Add Customer")
		fmt.Println("2. Display Customers")
		fmt.Println("3. Add Product")
		fmt.Println("4. Display Products")
		fmt.Println("5. Low Stock Products")
		fmt.Println("6. Create Order")
		fmt.Println("7. Display Order Details")
		fmt.Println("8. Exit")

		choice := readInt(reader, "Enter choice: ")

		switch choice {

		case 1:
			addCustomer(conn, reader)

		case 2:
			displayCustomers(conn)

		case 3:
			addProduct(conn, reader)

		case 4:
			displayProducts(conn)

		case 5:
			displayLowStock(conn)

		case 6:
			createOrder(conn, reader)

		case 7:
			displayOrders(conn)

		case 8:
			fmt.Println("Program ended.")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// ==================== CUSTOMER ====================

func addCustomer(conn *pgx.Conn, reader *bufio.Reader) {

	name := readString(reader, "Enter name: ")
	email := readString(reader, "Enter email: ")

	err := controller.AddCustomer(conn, model.Customer{
		Name:  name,
		Email: email,
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Customer added successfully.")
}

func displayCustomers(conn *pgx.Conn) {

	customers, err := controller.GetCustomers(conn)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("===== CUSTOMERS =====")
	fmt.Println("ID | Name | Email")

	if len(customers) == 0 {
		fmt.Println("No customers found.")
		return
	}

	for _, c := range customers {

		fmt.Printf(
			"%d | %s | %s\n",
			c.ID,
			c.Name,
			c.Email,
		)
	}
}

// ==================== PRODUCT ====================

func addProduct(conn *pgx.Conn, reader *bufio.Reader) {

	name := readString(reader, "Enter product name: ")
	price := readFloat(reader, "Enter price: ")
	stock := readInt(reader, "Enter stock: ")
	limit := readInt(reader, "Enter low stock limit: ")

	err := controller.AddProduct(conn, model.Product{
		Name:          name,
		Price:         price,
		Stock:         stock,
		LowStockLimit: limit,
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Product added successfully.")
}

func displayProducts(conn *pgx.Conn) {

	products, err := controller.GetProducts(conn)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("===== PRODUCTS =====")

	if len(products) == 0 {
		fmt.Println("No products found.")
		return
	}

	printProducts(products)
}

func displayLowStock(conn *pgx.Conn) {

	products, err := controller.GetLowStockProducts(conn)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("===== LOW STOCK PRODUCTS =====")

	if len(products) == 0 {
		fmt.Println("No low-stock products found.")
		return
	}

	printProducts(products)
}

func printProducts(products []model.Product) {

	fmt.Println("ID | Name | Price | Stock | Limit")

	for _, p := range products {

		fmt.Printf(
			"%d | %s | %.2f | %d | %d\n",
			p.ID,
			p.Name,
			p.Price,
			p.Stock,
			p.LowStockLimit,
		)
	}
}

// ==================== ORDER ====================

func createOrder(conn *pgx.Conn, reader *bufio.Reader) {

	customerID := readInt(reader, "Enter customer ID: ")

	count := readInt(reader, "How many products? ")

	if count <= 0 {
		fmt.Println("Number of products must be greater than 0.")
		return
	}

	var items []model.OrderItem

	for i := 0; i < count; i++ {

		fmt.Println()
		fmt.Println("Product", i+1)

		productID := readInt(reader, "Enter product ID: ")
		quantity := readInt(reader, "Enter quantity: ")

		if quantity <= 0 {
			fmt.Println("Quantity must be greater than 0.")
			return
		}

		items = append(items, model.OrderItem{
			ProductID: productID,
			Quantity:  quantity,
		})
	}

	orderID, err := controller.CreateOrder(
		conn,
		customerID,
		items,
	)

	if err != nil {

		fmt.Println("Order failed:", err)
		fmt.Println("Transaction rolled back.")

		return
	}

	fmt.Println()
	fmt.Println("Order created successfully.")
	fmt.Println("Order ID:", orderID)
}

func displayOrders(conn *pgx.Conn) {

	orders, err := controller.GetOrderDetails(conn)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("===== COMPLETE ORDER DETAILS =====")

	if len(orders) == 0 {
		fmt.Println("No orders found.")
		return
	}

	for _, o := range orders {

		fmt.Printf(
			"Order ID: %d | Customer: %s | Product: %s | Quantity: %d | Unit Price: %.2f | Line Total: %.2f | Order Total: %.2f\n",
			o.OrderID,
			o.CustomerName,
			o.ProductName,
			o.Quantity,
			o.UnitPrice,
			o.LineTotal,
			o.OrderTotal,
		)
	}
}

// ==================== INPUT FUNCTIONS ====================

func readString(reader *bufio.Reader, message string) string {

	for {

		fmt.Print(message)

		value, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		value = strings.TrimSpace(value)

		if value == "" {
			fmt.Println("Input cannot be empty.")
			continue
		}

		return value
	}
}

func readInt(reader *bufio.Reader, message string) int {

	for {

		value := readString(reader, message)

		number, err := strconv.Atoi(value)

		if err == nil {
			return number
		}

		fmt.Println("Please enter a valid number.")
	}
}

func readFloat(reader *bufio.Reader, message string) float64 {

	for {

		value := readString(reader, message)

		number, err := strconv.ParseFloat(value, 64)

		if err == nil {
			return number
		}

		fmt.Println("Please enter a valid price.")
	}
}
