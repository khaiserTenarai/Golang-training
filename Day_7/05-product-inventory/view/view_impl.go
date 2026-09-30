package view
import (
	"fmt"

	"product_inventory/model"
)

type ProductViewImpl struct {
}

func NewProductView() ProductView {

	return &ProductViewImpl{}
}

func (v *ProductViewImpl) ShowMenu() int {

	fmt.Println("\n========== Product Inventory System ==========")

	fmt.Println("1. Save Product")
	fmt.Println("2. Find Product")
	fmt.Println("3. Find All Products")
	fmt.Println("4. Update Product")
	fmt.Println("5. Delete Product")
	fmt.Println("6. Increase Stock")
	fmt.Println("7. Decrease Stock")
	fmt.Println("8. Low Stock Products")
	fmt.Println("9. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *ProductViewImpl) ReadProduct() model.Product {

	var product model.Product

	fmt.Println("\n---------- Enter Product ----------")

	fmt.Print("Enter Name: ")
	fmt.Scan(&product.Name)

	fmt.Print("Enter Price: ")
	fmt.Scan(&product.Price)

	fmt.Print("Enter Quantity: ")
	fmt.Scan(&product.Quantity)

	return product
}

func (v *ProductViewImpl) ReadProductForUpdate() model.Product {

	var product model.Product

	fmt.Println("\n---------- Update Product ----------")

	fmt.Print("Enter Product ID: ")
	fmt.Scan(&product.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&product.Name)

	fmt.Print("Enter Price: ")
	fmt.Scan(&product.Price)

	fmt.Print("Enter Quantity: ")
	fmt.Scan(&product.Quantity)

	return product
}

func (v *ProductViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter Product ID: ")
	fmt.Scan(&id)

	return id
}

func (v *ProductViewImpl) ReadStockQuantity() int {

	var quantity int

	fmt.Print("Enter stock quantity: ")
	fmt.Scan(&quantity)

	return quantity
}

func (v *ProductViewImpl) ReadLowStockLimit() int {

	var limit int

	fmt.Print("Enter low stock limit: ")
	fmt.Scan(&limit)

	return limit
}

func (v *ProductViewImpl) DisplayProduct(
	product model.Product,
) {

	fmt.Println("\n---------- Product ----------")

	fmt.Println("ID       :", product.ID)
	fmt.Println("Name     :", product.Name)
	fmt.Println("Price    :", product.Price)
	fmt.Println("Quantity :", product.Quantity)
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
			product.ID,
			product.Name,
			product.Price,
			product.Quantity,
		)
	}
}
