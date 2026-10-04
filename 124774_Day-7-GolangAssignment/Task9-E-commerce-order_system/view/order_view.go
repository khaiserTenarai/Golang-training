package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ecommerce-order/model"
)

type OrderView struct {
	reader *bufio.Reader
}

func NewOrderView() *OrderView {

	return &OrderView{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (v *OrderView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== E-COMMERCE ORDER SYSTEM ==========")
	fmt.Println("1. Create Customer")
	fmt.Println("2. Create Product")
	fmt.Println("3. Create Order")
	fmt.Println("4. Add Order Item")
	fmt.Println("5. View Order Details")
	fmt.Println("6. Exit")
	fmt.Println("=============================================")
}

func (v *OrderView) ReadInt(message string) int {

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

func (v *OrderView) ReadFloat(message string) float64 {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err == nil {
			return value
		}

		fmt.Println("Enter a valid price.")
	}
}

func (v *OrderView) ReadString(message string) string {

	fmt.Print(message)

	input, _ := v.reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func (v *OrderView) ReadCustomer() model.Customer {

	return model.Customer{
		Name:  v.ReadString("Enter customer name: "),
		Email: v.ReadString("Enter customer email: "),
	}
}

func (v *OrderView) ReadProduct() model.Product {

	return model.Product{
		Name:  v.ReadString("Enter product name: "),
		Price: v.ReadFloat("Enter product price: "),
	}
}

func (v *OrderView) ShowOrderDetails(
	details []model.OrderDetail,
) {

	for _, d := range details {

		fmt.Println("-----------------------------")
		fmt.Println("Order ID      :", d.OrderID)
		fmt.Println("Customer      :", d.CustomerName)
		fmt.Println("Product       :", d.ProductName)
		fmt.Println("Quantity      :", d.Quantity)
		fmt.Println("Price         :", d.Price)
		fmt.Println("-----------------------------")
	}
}

func (v *OrderView) ShowMessage(message string) {

	fmt.Println(message)
}
