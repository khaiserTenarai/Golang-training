package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"prodcut-inventory/model"
)

type ProductView struct {
	reader *bufio.Reader
}

func NewProductView() *ProductView {

	return &ProductView{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (v *ProductView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== PRODUCT INVENTORY ==========")
	fmt.Println("1. Add Product")
	fmt.Println("2. Find Product By ID")
	fmt.Println("3. Find All Products")
	fmt.Println("4. Update Product")
	fmt.Println("5. Delete Product")
	fmt.Println("6. Increase Stock")
	fmt.Println("7. Decrease Stock")
	fmt.Println("8. Low Stock Products")
	fmt.Println("9. Exit")
	fmt.Println("=======================================")
}

func (v *ProductView) ReadInt(
	message string,
) int {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err == nil {
			return value
		}

		fmt.Println("Enter a valid number.")
	}
}

func (v *ProductView) ReadFloat(
	message string,
) float64 {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err == nil {
			return value
		}

		fmt.Println("Enter a valid number.")
	}
}

func (v *ProductView) ReadString(
	message string,
) string {

	fmt.Print(message)

	input, _ := v.reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func (v *ProductView) ReadProduct() model.Product {

	return model.Product{
		Name: v.ReadString(
			"Enter Product Name: ",
		),

		Description: v.ReadString(
			"Enter Description: ",
		),

		Price: v.ReadFloat(
			"Enter Price: ",
		),

		Stock: v.ReadInt(
			"Enter Stock: ",
		),

		LowStockLimit: v.ReadInt(
			"Enter Low Stock Limit: ",
		),
	}
}

func (v *ProductView) ReadProductForUpdate() model.Product {

	return model.Product{
		ID: v.ReadInt(
			"Enter Product ID: ",
		),

		Name: v.ReadString(
			"Enter Product Name: ",
		),

		Description: v.ReadString(
			"Enter Description: ",
		),

		Price: v.ReadFloat(
			"Enter Price: ",
		),

		Stock: v.ReadInt(
			"Enter Stock: ",
		),

		LowStockLimit: v.ReadInt(
			"Enter Low Stock Limit: ",
		),
	}
}

func (v *ProductView) ShowProduct(
	product model.Product,
) {

	fmt.Println("--------------------------------")
	fmt.Println("ID             :", product.ID)
	fmt.Println("Name           :", product.Name)
	fmt.Println("Description    :", product.Description)
	fmt.Println("Price          :", product.Price)
	fmt.Println("Stock          :", product.Stock)
	fmt.Println("Low Stock Limit:", product.LowStockLimit)
	fmt.Println("--------------------------------")
}

func (v *ProductView) ShowProducts(
	products []model.Product,
) {

	if len(products) == 0 {
		fmt.Println("No products found.")
		return
	}

	for _, product := range products {
		v.ShowProduct(product)
	}
}

func (v *ProductView) ShowMessage(
	message string,
) {

	fmt.Println(message)
}
