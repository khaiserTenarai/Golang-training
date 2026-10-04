package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task10_order_inventory_transaction/models"
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
	fmt.Println("\n===== ORDER & INVENTORY TRANSACTION =====")
	fmt.Println("1. Add Product")
	fmt.Println("2. List Products")
	fmt.Println("3. Place Order (auto stock reduction)")
	fmt.Println("4. List Orders")
	fmt.Println("5. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
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

func (v *OrderView) GetOrderInput() (int, int) {
	fmt.Print("Enter Product ID: ")
	productID, _ := strconv.Atoi(v.readLine())
	fmt.Print("Enter Quantity: ")
	quantity, _ := strconv.Atoi(v.readLine())
	return productID, quantity
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

func (v *OrderView) ShowOrders(orders []models.Order) {
	if len(orders) == 0 {
		fmt.Println("\nNo orders found.")
		return
	}
	fmt.Printf("\n%-5s %-20s %-8s %-12s %-10s %-20s\n", "ID", "Product", "Qty", "Total", "Status", "Date")
	fmt.Println(strings.Repeat("-", 80))
	for _, o := range orders {
		fmt.Printf("%-5d %-20s %-8d %-12.2f %-10s %-20s\n",
			o.ID, o.ProductName, o.Quantity, o.TotalPrice, o.Status, o.CreatedAt.Format("2006-01-02 15:04"))
	}
}

func (v *OrderView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *OrderView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
