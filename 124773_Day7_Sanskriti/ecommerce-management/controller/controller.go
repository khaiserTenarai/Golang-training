package controller

import (
	"fmt"

	"ecommerce-management/model"
	"ecommerce-management/service"
)

type Controller struct {
	Service *service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{
		Service: s,
	}
}

// Add Product
func (c *Controller) AddProduct() {

	var product model.Product

	fmt.Print("Enter Product Name: ")
	fmt.Scanln(&product.Name)

	fmt.Print("Enter Price: ")
	fmt.Scanln(&product.Price)

	fmt.Print("Enter Stock: ")
	fmt.Scanln(&product.Stock)

	fmt.Print("Enter Low Stock Limit: ")
	fmt.Scanln(&product.LowStockLimit)

	err := c.Service.AddProduct(&product)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Product added successfully")
}

// Display Products
func (c *Controller) DisplayProducts() {

	products, err := c.Service.GetProducts()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No products found")
		return
	}

	fmt.Println("\n------ PRODUCTS ------")

	for _, product := range products {

		fmt.Println("ID:", product.ID)
		fmt.Println("Name:", product.Name)
		fmt.Println("Price:", product.Price)
		fmt.Println("Stock:", product.Stock)
		fmt.Println("Low Stock Limit:", product.LowStockLimit)
		fmt.Println("----------------------")
	}
}

// Update Product
func (c *Controller) UpdateProduct() {

	var product model.Product

	fmt.Print("Enter Product ID: ")
	fmt.Scanln(&product.ID)

	fmt.Print("Enter New Name: ")
	fmt.Scanln(&product.Name)

	fmt.Print("Enter New Price: ")
	fmt.Scanln(&product.Price)

	fmt.Print("Enter New Stock: ")
	fmt.Scanln(&product.Stock)

	fmt.Print("Enter New Low Stock Limit: ")
	fmt.Scanln(&product.LowStockLimit)

	err := c.Service.UpdateProduct(&product)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Product updated successfully")
}

// Delete Product
func (c *Controller) DeleteProduct() {

	var id int

	fmt.Print("Enter Product ID: ")
	fmt.Scanln(&id)

	err := c.Service.DeleteProduct(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Product deleted successfully")
}

// Increase Stock
func (c *Controller) IncreaseStock() {

	var id int
	var quantity int

	fmt.Print("Enter Product ID: ")
	fmt.Scanln(&id)

	fmt.Print("Enter quantity to increase: ")
	fmt.Scanln(&quantity)

	err := c.Service.IncreaseStock(id, quantity)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Stock increased successfully")
}

// Decrease Stock
func (c *Controller) DecreaseStock() {

	var id int
	var quantity int

	fmt.Print("Enter Product ID: ")
	fmt.Scanln(&id)

	fmt.Print("Enter quantity to decrease: ")
	fmt.Scanln(&quantity)

	err := c.Service.DecreaseStock(id, quantity)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Stock decreased successfully")
}

// Low Stock Products
func (c *Controller) LowStockProducts() {

	products, err := c.Service.GetLowStockProducts()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("No low-stock products")
		return
	}

	fmt.Println("\n------ LOW STOCK PRODUCTS ------")

	for _, product := range products {

		fmt.Println("ID:", product.ID)
		fmt.Println("Name:", product.Name)
		fmt.Println("Stock:", product.Stock)
		fmt.Println("Low Stock Limit:", product.LowStockLimit)
		fmt.Println("-------------------------------")
	}
}
