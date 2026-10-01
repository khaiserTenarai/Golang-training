package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/employee-management/model"
	"example.com/employee-management/repository"
	"example.com/employee-management/service"
)

type ProductController struct {
	service service.ProductService
	reader  *bufio.Reader
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (c *ProductController) Start() {
	for {
		c.showMenu()
		choice := c.readInt("Enter your choice: ")
		switch choice {
		case 1:
			c.addProduct()
		case 2:
			c.getProduct()
		case 3:
			c.getAllProducts()
		case 4:
			c.updateProduct()
		case 5:
			c.increaseStock()
		case 6:
			c.decreaseStock()
		case 7:
			c.getLowStockProducts()
		case 8:
			c.deleteProduct()
		case 9:
			fmt.Println()
			fmt.Println("Thank you for using Product Inventory System.")
			return
		default:
			fmt.Println("Invalid choice. Please select 1 to 9.")
		}
		c.pause()
	}
}

func (c *ProductController) showMenu() {
	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("       PRODUCT INVENTORY MANAGEMENT")
	fmt.Println("==============================================")
	fmt.Println("1. Add Product")
	fmt.Println("2. Find Product")
	fmt.Println("3. List All Products")
	fmt.Println("4. Update Product")
	fmt.Println("5. Increase Stock")
	fmt.Println("6. Decrease Stock")
	fmt.Println("7. Search Low-Stock Products")
	fmt.Println("8. Delete Product")
	fmt.Println("9. Exit")
	fmt.Println("==============================================")
}

func (c *ProductController) addProduct() {
	fmt.Println()
	fmt.Println("---------- ADD PRODUCT ----------")
	var product model.Product
	product.Name = c.readString("Enter Name: ")
	product.Category = c.readString("Enter Category: ")
	product.Price = c.readFloat("Enter Price: ")
	product.StockQuantity = c.readInt("Enter Stock Quantity: ")
	product.ReorderLevel = c.readInt("Enter Reorder Level: ")

	err := c.service.AddProduct(product)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println()
	fmt.Println("Product added successfully.")
}

func (c *ProductController) getProduct() {
	fmt.Println()
	fmt.Println("---------- FIND PRODUCT ----------")
	id := c.readInt64("Enter Product ID: ")
	product, err := c.service.GetProduct(id)
	if errors.Is(err, repository.ErrProductNotFound) {
		fmt.Println("Product not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	product.Display()
}

func (c *ProductController) getAllProducts() {
	fmt.Println()
	fmt.Println("---------- ALL PRODUCTS ----------")
	products, err := c.service.GetAllProducts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(products) == 0 {
		fmt.Println("No products found.")
		return
	}
	c.printProductTable(products)
}

func (c *ProductController) updateProduct() {
	fmt.Println()
	fmt.Println("---------- UPDATE PRODUCT ----------")
	id := c.readInt64("Enter Product ID: ")
	product, err := c.service.GetProduct(id)
	if errors.Is(err, repository.ErrProductNotFound) {
		fmt.Println("Product not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("Current product:")
	product.Display()

	fmt.Println()
	fmt.Println("Enter new values.")
	product.Name = c.readString("Enter Name: ")
	product.Category = c.readString("Enter Category: ")
	product.Price = c.readFloat("Enter Price: ")
	product.StockQuantity = c.readInt("Enter Stock Quantity: ")
	product.ReorderLevel = c.readInt("Enter Reorder Level: ")

	err = c.service.UpdateProduct(product)
	if errors.Is(err, repository.ErrProductNotFound) {
		fmt.Println("Product not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Product updated successfully.")
}

func (c *ProductController) increaseStock() {
	fmt.Println()
	fmt.Println("---------- INCREASE STOCK ----------")
	id := c.readInt64("Enter Product ID: ")
	amount := c.readInt("Enter Quantity to Add: ")

	err := c.service.IncreaseStock(id, amount)
	if errors.Is(err, repository.ErrProductNotFound) {
		fmt.Println("Product not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Stock increased successfully.")
}

func (c *ProductController) decreaseStock() {
	fmt.Println()
	fmt.Println("---------- DECREASE STOCK ----------")
	id := c.readInt64("Enter Product ID: ")
	amount := c.readInt("Enter Quantity to Reduce: ")

	err := c.service.DecreaseStock(id, amount)
	if errors.Is(err, repository.ErrProductNotFound) {
		fmt.Println("Product not found.")
		return
	}
	if errors.Is(err, repository.ErrInsufficientStock) {
		fmt.Println("Error: Insufficient stock available.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Stock decreased successfully.")
}

func (c *ProductController) getLowStockProducts() {
	fmt.Println()
	fmt.Println("---------- LOW STOCK PRODUCTS ----------")
	products, err := c.service.GetLowStockProducts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(products) == 0 {
		fmt.Println("No low-stock products found.")
		return
	}
	c.printProductTable(products)
}

func (c *ProductController) deleteProduct() {
	fmt.Println()
	fmt.Println("---------- DELETE PRODUCT ----------")
	id := c.readInt64("Enter Product ID: ")
	product, err := c.service.GetProduct(id)
	if errors.Is(err, repository.ErrProductNotFound) {
		fmt.Println("Product not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	product.Display()

	confirmation := c.readString("Are you sure you want to delete? (y/n): ")
	switch strings.ToLower(confirmation) {
	case "y", "yes":
		err = c.service.DeleteProduct(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Product deleted successfully.")
	case "n", "no":
		fmt.Println("Delete operation cancelled.")
	default:
		fmt.Println("Invalid response. Delete operation cancelled.")
	}
}

func (c *ProductController) printProductTable(products []model.Product) {
	fmt.Println()
	fmt.Printf(
		"%-5s %-25s %-15s %-10s %-8s %-12s\n",
		"ID",
		"NAME",
		"CATEGORY",
		"PRICE",
		"STOCK",
		"REORDER LBL",
	)
	fmt.Println(strings.Repeat("-", 80))
	for _, p := range products {
		fmt.Printf(
			"%-5d %-25s %-15s %-10.2f %-8d %-12d\n",
			p.ID,
			p.Name,
			p.Category,
			p.Price,
			p.StockQuantity,
			p.ReorderLevel,
		)
	}
}

func (c *ProductController) readString(message string) string {
	for {
		fmt.Print(message)
		value, err := c.reader.ReadString('\n')
		if err != nil {
			continue
		}
		value = strings.TrimSpace(value)
		if value != "" {
			return value
		}
		fmt.Println("Value cannot be empty.")
	}
}

func (c *ProductController) readInt(message string) int {
	for {
		value := c.readString(message)
		number, err := strconv.Atoi(value)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid integer.")
	}
}

func (c *ProductController) readInt64(message string) int64 {
	for {
		value := c.readString(message)
		number, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid number.")
	}
}

func (c *ProductController) readFloat(message string) float64 {
	for {
		value := c.readString(message)
		number, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid decimal number.")
	}
}

func (c *ProductController) pause() {
	fmt.Println()
	fmt.Print("Press ENTER to continue...")
	_, _ = c.reader.ReadString('\n')
}