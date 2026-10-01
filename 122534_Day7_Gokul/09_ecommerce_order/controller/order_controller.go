package controller

import (
	"bufio"
	"context"
	"example.com/q9-ecommerce-order/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OrderController struct {
	service service.OrderService
	reader  *bufio.Reader
}

func NewOrderController(s service.OrderService) *OrderController {
	return &OrderController{service: s, reader: bufio.NewReader(os.Stdin)}
}
func (c *OrderController) Start() {
	for {
		fmt.Println("\n1. Add Customer\n2. Add Product\n3. Create Order\n4. Show Order Details\n5. Exit")
		ch := c.readInt("Choice: ")
		ctx := context.Background()
		switch ch {
		case 1:
			c.done(c.service.AddCustomer(ctx, c.readString("Customer name: "), c.readString("Email: ")))
		case 2:
			c.done(c.service.AddProduct(ctx, c.readString("Product name: "), c.readFloat("Price: ")))
		case 3:
			c.done(c.service.CreateOrder(ctx, c.readInt64("Customer ID: "), c.readInt64("Product ID: "), c.readInt("Quantity: ")))
		case 4:
			c.done(c.service.ShowOrders(ctx))
		case 5:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
func (c *OrderController) done(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Operation completed.")
	}
}
func (c *OrderController) readString(p string) string {
	fmt.Print(p)
	v, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(v)
}
func (c *OrderController) readInt(p string) int { v, _ := strconv.Atoi(c.readString(p)); return v }
func (c *OrderController) readInt64(p string) int64 {
	v, _ := strconv.ParseInt(c.readString(p), 10, 64)
	return v
}
func (c *OrderController) readFloat(p string) float64 {
	v, _ := strconv.ParseFloat(c.readString(p), 64)
	return v
}
