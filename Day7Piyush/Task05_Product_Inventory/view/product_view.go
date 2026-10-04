package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task05_product_inventory/models"
)

type ProductView struct {
	scanner *bufio.Scanner
}

func NewProductView() *ProductView {
	return &ProductView{scanner: bufio.NewScanner(os.Stdin)}
}

func (v *ProductView) readLine() string {
	v.scanner.Scan()
	return strings.TrimSpace(v.scanner.Text())
}

func (v *ProductView) ShowMenu() int {
	fmt.Println("\n===== PRODUCT INVENTORY =====")
	fmt.Println("1. Add Product")
	fmt.Println("2. List All Products")
	fmt.Println("3. Update Product")
	fmt.Println("4. Delete Product")
	fmt.Println("5. Increase Stock")
	fmt.Println("6. Decrease Stock")
	fmt.Println("7. Low-Stock Products")
	fmt.Println("8. Search Product by Name")
	fmt.Println("9. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
}

func (v *ProductView) GetProductInput() models.Product {
	var p models.Product
	fmt.Print("Enter Product Name: ")
	p.Name = v.readLine()
	fmt.Print("Enter Category: ")
	p.Category = v.readLine()
	fmt.Print("Enter Price: ")
	p.Price, _ = strconv.ParseFloat(v.readLine(), 64)
	fmt.Print("Enter Stock: ")
	p.Stock, _ = strconv.Atoi(v.readLine())
	return p
}

func (v *ProductView) GetID() int {
	fmt.Print("Enter Product ID: ")
	id, _ := strconv.Atoi(v.readLine())
	return id
}

func (v *ProductView) GetQuantity() int {
	fmt.Print("Enter Quantity: ")
	qty, _ := strconv.Atoi(v.readLine())
	return qty
}

func (v *ProductView) GetThreshold() int {
	fmt.Print("Enter low-stock threshold: ")
	t, _ := strconv.Atoi(v.readLine())
	return t
}

func (v *ProductView) GetSearchName() string {
	fmt.Print("Enter product name to search: ")
	return v.readLine()
}

func (v *ProductView) GetUpdateInput(existing models.Product) models.Product {
	fmt.Printf("Name (%s) - new value (Enter to keep): ", existing.Name)
	if val := v.readLine(); val != "" {
		existing.Name = val
	}
	fmt.Printf("Category (%s) - new value (Enter to keep): ", existing.Category)
	if val := v.readLine(); val != "" {
		existing.Category = val
	}
	fmt.Printf("Price (%.2f) - new value (Enter to keep): ", existing.Price)
	if val := v.readLine(); val != "" {
		existing.Price, _ = strconv.ParseFloat(val, 64)
	}
	fmt.Printf("Stock (%d) - new value (Enter to keep): ", existing.Stock)
	if val := v.readLine(); val != "" {
		existing.Stock, _ = strconv.Atoi(val)
	}
	return existing
}

func (v *ProductView) ShowProducts(products []models.Product) {
	if len(products) == 0 {
		fmt.Println("\nNo products found.")
		return
	}
	fmt.Printf("\n%-5s %-25s %-15s %-10s %-8s\n", "ID", "Name", "Category", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 68))
	for _, p := range products {
		fmt.Printf("%-5d %-25s %-15s %-10.2f %-8d\n", p.ID, p.Name, p.Category, p.Price, p.Stock)
	}
}

func (v *ProductView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *ProductView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
