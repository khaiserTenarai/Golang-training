package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"product-inventory/controller"
	"product-inventory/model"
)

type ProductViewImpl struct {
	controller controller.ProductController
	reader     *bufio.Reader
}

func NewProductView(controller controller.ProductController) ProductView {
	return &ProductViewImpl{
		controller: controller,
		reader:     bufio.NewReader(os.Stdin),
	}
}

func (v *ProductViewImpl) Start() {
	for {
		fmt.Println("\n----- PRODUCT INVENTORY -----")
		fmt.Println("1. Create Product")
		fmt.Println("2. Get Product")
		fmt.Println("3. Get All Products")
		fmt.Println("4. Update Product")
		fmt.Println("5. Delete Product")
		fmt.Println("6. Increase Stock")
		fmt.Println("7. Decrease Stock")
		fmt.Println("8. Low Stock Products")
		fmt.Println("9. Exit")

		fmt.Print("Enter choice: ")
		choice := v.readInt()

		switch choice {
		case 1:
			v.CreateProduct()
		case 2:
			v.GetProduct()
		case 3:
			v.GetAllProducts()
		case 4:
			v.UpdateProduct()
		case 5:
			v.DeleteProduct()
		case 6:
			v.IncreaseStock()
		case 7:
			v.DecreaseStock()
		case 8:
			v.GetLowStockProducts()
		case 9:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (v *ProductViewImpl) CreateProduct() {
	var product model.Product

	fmt.Print("Enter Product Name: ")
	product.Name = v.readString()

	fmt.Print("Enter Price: ")
	product.Price = v.readFloat()

	fmt.Print("Enter Stock: ")
	product.Stock = v.readInt()

	fmt.Print("Enter Low Stock Threshold: ")
	product.LowStockThreshold = v.readInt()

	err := v.controller.CreateProduct(product)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Product created successfully")
}

func (v *ProductViewImpl) GetProduct() {
	fmt.Print("Enter Product ID: ")
	id := v.readInt()

	product, err := v.controller.GetProduct(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayProduct(product)
}

func (v *ProductViewImpl) GetAllProducts() {
	products, err := v.controller.GetAllProducts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayProducts(products)
}

func (v *ProductViewImpl) UpdateProduct() {
	var product model.Product

	fmt.Print("Enter Product ID: ")
	product.ID = v.readInt()

	fmt.Print("Enter Product Name: ")
	product.Name = v.readString()

	fmt.Print("Enter Price: ")
	product.Price = v.readFloat()

	fmt.Print("Enter Stock: ")
	product.Stock = v.readInt()

	fmt.Print("Enter Low Stock Threshold: ")
	product.LowStockThreshold = v.readInt()

	err := v.controller.UpdateProduct(product)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Product updated successfully")
}

func (v *ProductViewImpl) DeleteProduct() {
	fmt.Print("Enter Product ID: ")
	id := v.readInt()

	err := v.controller.DeleteProduct(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Product deleted successfully")
}

func (v *ProductViewImpl) IncreaseStock() {
	fmt.Print("Enter Product ID: ")
	id := v.readInt()

	fmt.Print("Enter Quantity: ")
	quantity := v.readInt()

	err := v.controller.IncreaseStock(id, quantity)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Stock increased successfully")
}

func (v *ProductViewImpl) DecreaseStock() {
	fmt.Print("Enter Product ID: ")
	id := v.readInt()

	fmt.Print("Enter Quantity: ")
	quantity := v.readInt()

	err := v.controller.DecreaseStock(id, quantity)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Stock decreased successfully")
}

func (v *ProductViewImpl) GetLowStockProducts() {
	products, err := v.controller.GetLowStockProducts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayProducts(products)
}

func (v *ProductViewImpl) DisplayProduct(product *model.Product) {
	fmt.Println("\n----- PRODUCT DETAILS -----")
	fmt.Println("ID                 :", product.ID)
	fmt.Println("Name               :", product.Name)
	fmt.Println("Price              :", product.Price)
	fmt.Println("Stock              :", product.Stock)
	fmt.Println("Low Stock Threshold:", product.LowStockThreshold)
}

func (v *ProductViewImpl) DisplayProducts(products []model.Product) {
	fmt.Println("\n----- PRODUCT LIST -----")

	if len(products) == 0 {
		fmt.Println("No products found")
		return
	}

	for _, product := range products {
		fmt.Println("----------------------------")
		fmt.Println("ID                 :", product.ID)
		fmt.Println("Name               :", product.Name)
		fmt.Println("Price              :", product.Price)
		fmt.Println("Stock              :", product.Stock)
		fmt.Println("Low Stock Threshold:", product.LowStockThreshold)
	}
}

func (v *ProductViewImpl) readString() string {
	value, _ := v.reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func (v *ProductViewImpl) readInt() int {
	value := v.readString()
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return number
}

func (v *ProductViewImpl) readFloat() float64 {
	value := v.readString()
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return number
}
