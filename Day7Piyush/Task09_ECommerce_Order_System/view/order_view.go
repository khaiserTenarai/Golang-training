package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task09_ecommerce_order_system/models"
)

type OrderView struct {
	scanner *bufio.Scanner
}

func NewOrderView() *OrderView {
	return &OrderView{scanner: bufio.NewScanner(os.Stdin)}
}

func (v *OrderView) readLine() string {
	v.scanner.Scan()
	return strings.TrimSpace(v.scanner.Text())
}

func (v *OrderView) ShowMenu() int {
	fmt.Println("\n===== E-COMMERCE ORDER SYSTEM =====")
	fmt.Println("1. Add Customer")
	fmt.Println("2. Add Product")
	fmt.Println("3. List Customers")
	fmt.Println("4. List Products")
	fmt.Println("5. Create Order")
	fmt.Println("6. View Order Details (JOINs)")
	fmt.Println("7. List All Orders")
	fmt.Println("8. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
}

func (v *OrderView) GetCustomerInput() models.Customer {
	var c models.Customer
	fmt.Print("Enter Customer Name: ")
	c.Name = v.readLine()
	fmt.Print("Enter Customer Email: ")
	c.Email = v.readLine()
	return c
}

func (v *OrderView) GetProductInput() models.Product {
	var p models.Product
	fmt.Print("Enter Product Name: ")
	p.Name = v.readLine()
	fmt.Print("Enter Price: ")
	p.Price, _ = strconv.ParseFloat(v.readLine(), 64)
	fmt.Print("Enter Stock: ")
	p.Stock, _ = strconv.Atoi(v.readLine())
	return p
}

func (v *OrderView) GetID(label string) int {
	fmt.Printf("Enter %s ID: ", label)
	id, _ := strconv.Atoi(v.readLine())
	return id
}

func (v *OrderView) GetOrderItems() (int, []models.OrderItem) {
	fmt.Print("Enter Customer ID: ")
	custID, _ := strconv.Atoi(v.readLine())
	fmt.Print("How many items? ")
	count, _ := strconv.Atoi(v.readLine())
	var items []models.OrderItem
	for i := 0; i < count; i++ {
		fmt.Printf("\n--- Item %d ---\n", i+1)
		var item models.OrderItem
		fmt.Print("Product ID: ")
		item.ProductID, _ = strconv.Atoi(v.readLine())
		fmt.Print("Quantity: ")
		item.Quantity, _ = strconv.Atoi(v.readLine())
		fmt.Print("Unit Price: ")
		item.UnitPrice, _ = strconv.ParseFloat(v.readLine(), 64)
		items = append(items, item)
	}
	return custID, items
}

func (v *OrderView) ShowCustomers(customers []models.Customer) {
	if len(customers) == 0 {
		fmt.Println("\nNo customers found.")
		return
	}
	fmt.Printf("\n%-5s %-20s %-30s\n", "ID", "Name", "Email")
	fmt.Println(strings.Repeat("-", 58))
	for _, c := range customers {
		fmt.Printf("%-5d %-20s %-30s\n", c.ID, c.Name, c.Email)
	}
}

func (v *OrderView) ShowProducts(products []models.Product) {
	if len(products) == 0 {
		fmt.Println("\nNo products found.")
		return
	}
	fmt.Printf("\n%-5s %-25s %-10s %-8s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 52))
	for _, p := range products {
		fmt.Printf("%-5d %-25s %-10.2f %-8d\n", p.ID, p.Name, p.Price, p.Stock)
	}
}

func (v *OrderView) ShowOrderDetails(order models.Order) {
	fmt.Printf("\n===== ORDER #%d =====\n", order.ID)
	fmt.Printf("Customer : %s (ID: %d)\n", order.CustomerName, order.CustomerID)
	fmt.Printf("Status   : %s\n", order.Status)
	fmt.Printf("Date     : %s\n", order.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Println("\nItems:")
	fmt.Printf("  %-5s %-25s %-8s %-12s %-12s\n", "ID", "Product", "Qty", "Unit Price", "Subtotal")
	fmt.Println("  " + strings.Repeat("-", 65))
	for _, item := range order.Items {
		fmt.Printf("  %-5d %-25s %-8d %-12.2f %-12.2f\n",
			item.ID, item.ProductName, item.Quantity, item.UnitPrice, item.Subtotal)
	}
	fmt.Printf("\nTotal Amount: %.2f\n", order.TotalAmount)
}

func (v *OrderView) ShowOrders(orders []models.Order) {
	if len(orders) == 0 {
		fmt.Println("\nNo orders found.")
		return
	}
	fmt.Printf("\n%-5s %-20s %-12s %-10s %-20s\n", "ID", "Customer", "Total", "Status", "Date")
	fmt.Println(strings.Repeat("-", 72))
	for _, o := range orders {
		fmt.Printf("%-5d %-20s %-12.2f %-10s %-20s\n",
			o.ID, o.CustomerName, o.TotalAmount, o.Status, o.CreatedAt.Format("2006-01-02 15:04"))
	}
}

func (v *OrderView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *OrderView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
