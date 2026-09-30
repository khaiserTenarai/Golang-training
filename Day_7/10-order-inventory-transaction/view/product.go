package view
import (
	"fmt"

	"order_inventory/model"
)

type ProductView interface {
	ShowMenu() int
	ReadProduct() model.Product
	ReadProductForUpdate() model.Product
	ReadID() int
	DisplayProduct(product model.Product)
	DisplayProducts(products []model.Product)
}

type ProductViewImpl struct {
}

func NewProductView() ProductView {
	return &ProductViewImpl{}
}

func (v *ProductViewImpl) ShowMenu() int {

	fmt.Println("\n========== Product Management ==========")
	fmt.Println("1. Add Product")
	fmt.Println("2. Find Product")
	fmt.Println("3. Find All Products")
	fmt.Println("4. Update Product")
	fmt.Println("5. Back")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *ProductViewImpl) ReadProduct() model.Product {

	var product model.Product

	fmt.Println("\n---------- Add Product ----------")

	fmt.Print("Enter Name: ")
	fmt.Scan(&product.Name)

	fmt.Print("Enter Price: ")
	fmt.Scan(&product.Price)

	fmt.Print("Enter Stock: ")
	fmt.Scan(&product.Stock)

	return product
}

func (v *ProductViewImpl) ReadProductForUpdate() model.Product {

	var product model.Product

	fmt.Println("\n---------- Update Product ----------")

	fmt.Print("Enter ID: ")
	fmt.Scan(&product.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&product.Name)

	fmt.Print("Enter Price: ")
	fmt.Scan(&product.Price)

	fmt.Print("Enter Stock: ")
	fmt.Scan(&product.Stock)

	return product
}

func (v *ProductViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter Product ID: ")
	fmt.Scan(&id)

	return id
}

func (v *ProductViewImpl) DisplayProduct(
	product model.Product,
) {

	fmt.Println("\n---------- Product ----------")
	fmt.Println("ID    :", product.ID)
	fmt.Println("Name  :", product.Name)
	fmt.Println("Price :", product.Price)
	fmt.Println("Stock :", product.Stock)
}

func (v *ProductViewImpl) DisplayProducts(
	products []model.Product,
) {

	if len(products) == 0 {
		fmt.Println("No products found.")
		return
	}

	fmt.Println("\n---------- Products ----------")

	for _, product := range products {

		fmt.Println(
			"ID:",
			product.ID,
			"| Name:",
			product.Name,
			"| Price:",
			product.Price,
			"| Stock:",
			product.Stock,
		)
	}
}
