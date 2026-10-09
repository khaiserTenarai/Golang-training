package controller

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"product-inventory/model"
	"product-inventory/service"
	"product-inventory/view"
)

type ProductController struct {
	service service.ProductService
	reader  *bufio.Reader
}

func NewProductController(
	service service.ProductService,
) *ProductController {

	return &ProductController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (c *ProductController) Start() {

	for {

		fmt.Println()
		fmt.Println("======================================")
		fmt.Println("       PRODUCT INVENTORY SYSTEM")
		fmt.Println("======================================")
		fmt.Println("1. Create Product")
		fmt.Println("2. Get Product")
		fmt.Println("3. Get All Products")
		fmt.Println("4. Update Product")
		fmt.Println("5. Delete Product")
		fmt.Println("6. Increase Stock")
		fmt.Println("7. Decrease Stock")
		fmt.Println("8. Low Stock Products")
		fmt.Println("9. Exit")
		fmt.Println("======================================")

		choice := c.readString(
			"Enter choice: ",
		)

		switch choice {
		case "1":
			c.createProduct()

		case "2":
			c.getProduct()

		case "3":
			c.getAllProducts()

		case "4":
			c.updateProduct()

		case "5":
			c.deleteProduct()

		case "6":
			c.increaseStock()

		case "7":
			c.decreaseStock()

		case "8":
			c.lowStockProducts()

		case "9":
			fmt.Println(
				"Exiting Product Inventory System...",
			)
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *ProductController) createProduct() {

	fmt.Println()
	fmt.Println("========== CREATE PRODUCT ==========")

	req := view.CreateProductRequest{
		Name: c.readString("Name: "),
		Description: c.readString(
			"Description: ",
		),
		Price: c.readFloat("Price: "),
		StockQuantity: c.readInt(
			"Initial Stock: ",
		),
		LowStockThreshold: c.readInt(
			"Low Stock Threshold: ",
		),
	}

	response, err :=
		c.service.CreateProduct(
			context.Background(),
			req,
		)

	if err != nil {
		fmt.Println(
			"Create failed:",
			err,
		)
		return
	}

	c.printProduct(response.Product)
}

func (c *ProductController) getProduct() {

	fmt.Println()
	fmt.Println("========== GET PRODUCT ==========")

	id := c.readInt64(
		"Product ID: ",
	)

	response, err :=
		c.service.GetProduct(
			context.Background(),
			id,
		)

	if err != nil {
		fmt.Println(
			"Get failed:",
			err,
		)
		return
	}

	c.printProduct(response.Product)
}

func (c *ProductController) getAllProducts() {

	fmt.Println()
	fmt.Println("========== ALL PRODUCTS ==========")

	response, err :=
		c.service.GetAllProducts(
			context.Background(),
		)

	if err != nil {
		fmt.Println(
			"Get products failed:",
			err,
		)
		return
	}

	c.printProducts(response.Products)
}

func (c *ProductController) updateProduct() {

	fmt.Println()
	fmt.Println("========== UPDATE PRODUCT ==========")

	req := view.UpdateProductRequest{
		ID: c.readInt64(
			"Product ID: ",
		),
		Name: c.readString(
			"Name: ",
		),
		Description: c.readString(
			"Description: ",
		),
		Price: c.readFloat(
			"Price: ",
		),
		LowStockThreshold: c.readInt(
			"Low Stock Threshold: ",
		),
	}

	response, err :=
		c.service.UpdateProduct(
			context.Background(),
			req,
		)

	if err != nil {
		fmt.Println(
			"Update failed:",
			err,
		)
		return
	}

	c.printProduct(response.Product)
}

func (c *ProductController) deleteProduct() {

	fmt.Println()
	fmt.Println("========== DELETE PRODUCT ==========")

	id := c.readInt64(
		"Product ID: ",
	)

	err :=
		c.service.DeleteProduct(
			context.Background(),
			id,
		)

	if err != nil {
		fmt.Println(
			"Delete failed:",
			err,
		)
		return
	}

	fmt.Println(
		"Product deleted successfully.",
	)
}

func (c *ProductController) increaseStock() {

	fmt.Println()
	fmt.Println("========== INCREASE STOCK ==========")

	req := view.StockRequest{
		ProductID: c.readInt64(
			"Product ID: ",
		),
		Quantity: c.readInt(
			"Quantity to add: ",
		),
	}

	response, err :=
		c.service.IncreaseStock(
			context.Background(),
			req,
		)

	if err != nil {
		fmt.Println(
			"Stock increase failed:",
			err,
		)
		return
	}

	fmt.Printf(
		"Stock increased successfully. New stock: %d\n",
		response.Product.StockQuantity,
	)
}

func (c *ProductController) decreaseStock() {

	fmt.Println()
	fmt.Println("========== DECREASE STOCK ==========")

	req := view.StockRequest{
		ProductID: c.readInt64(
			"Product ID: ",
		),
		Quantity: c.readInt(
			"Quantity to remove: ",
		),
	}

	response, err :=
		c.service.DecreaseStock(
			context.Background(),
			req,
		)

	if err != nil {
		fmt.Println(
			"Stock decrease failed:",
			err,
		)
		return
	}

	fmt.Printf(
		"Stock decreased successfully. New stock: %d\n",
		response.Product.StockQuantity,
	)
}

func (c *ProductController) lowStockProducts() {

	fmt.Println()
	fmt.Println("========== LOW STOCK PRODUCTS ==========")

	response, err :=
		c.service.GetLowStockProducts(
			context.Background(),
		)

	if err != nil {
		fmt.Println(
			"Low stock search failed:",
			err,
		)
		return
	}

	if len(response.Products) == 0 {

		fmt.Println(
			"No low-stock products found.",
		)

		return
	}

	c.printProducts(response.Products)
}

func (c *ProductController) printProduct(
	product model.Product,
) {

	fmt.Println()
	fmt.Println("--------------------------------------")
	fmt.Printf("ID              : %d\n", product.ID)
	fmt.Printf("Name            : %s\n", product.Name)
	fmt.Printf(
		"Description     : %s\n",
		product.Description,
	)
	fmt.Printf(
		"Price           : %.2f\n",
		product.Price,
	)
	fmt.Printf(
		"Stock           : %d\n",
		product.StockQuantity,
	)
	fmt.Printf(
		"Low Stock Level : %d\n",
		product.LowStockThreshold,
	)
	fmt.Printf(
		"Created At      : %s\n",
		product.CreatedAt.Format(
			"2006-01-02 15:04:05",
		),
	)
	fmt.Printf(
		"Updated At      : %s\n",
		product.UpdatedAt.Format(
			"2006-01-02 15:04:05",
		),
	)
	fmt.Println("--------------------------------------")
}

func (c *ProductController) printProducts(
	products []model.Product,
) {

	fmt.Println()

	fmt.Printf(
		"%-5s %-20s %-12s %-10s %-12s\n",
		"ID",
		"NAME",
		"PRICE",
		"STOCK",
		"LOW LEVEL",
	)

	fmt.Println(
		"--------------------------------------------------------------",
	)

	for _, product := range products {

		fmt.Printf(
			"%-5d %-20s %-12.2f %-10d %-12d\n",
			product.ID,
			product.Name,
			product.Price,
			product.StockQuantity,
			product.LowStockThreshold,
		)
	}
}

func (c *ProductController) readString(
	message string,
) string {

	fmt.Print(message)

	value, err :=
		c.reader.ReadString('\n')

	if err != nil {
		return ""
	}

	return strings.TrimSpace(value)
}

func (c *ProductController) readInt(
	message string,
) int {

	value := c.readString(message)

	result, err :=
		strconv.Atoi(value)

	if err != nil {
		return 0
	}

	return result
}

func (c *ProductController) readInt64(
	message string,
) int64 {

	value := c.readString(message)

	result, err :=
		strconv.ParseInt(
			value,
			10,
			64,
		)

	if err != nil {
		return 0
	}

	return result
}

func (c *ProductController) readFloat(
	message string,
) float64 {

	value := c.readString(message)

	result, err :=
		strconv.ParseFloat(
			value,
			64,
		)

	if err != nil {
		return 0
	}

	return result
}
