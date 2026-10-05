package controller

import (
	"bufio"
	"context"
	"example.com/q10-order-inventory/service"
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
		fmt.Println("\n1. Add Product\n2. View Products\n3. Create Order\n4. Exit")
		ch := c.i("Choice: ")
		ctx := context.Background()
		switch ch {
		case 1:
			c.done(c.service.AddProduct(ctx, c.s("Name: "), c.f("Price: "), c.i("Stock: ")))
		case 2:
			c.done(c.service.ListProducts(ctx))
		case 3:
			c.done(c.service.CreateOrder(ctx, c.i64("Product ID: "), c.i("Quantity: ")))
		case 4:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
func (c *OrderController) done(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println("Operation completed.")
	}
}
func (c *OrderController) s(p string) string {
	fmt.Print(p)
	v, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(v)
}
func (c *OrderController) i(p string) int     { v, _ := strconv.Atoi(c.s(p)); return v }
func (c *OrderController) i64(p string) int64 { v, _ := strconv.ParseInt(c.s(p), 10, 64); return v }
func (c *OrderController) f(p string) float64 { v, _ := strconv.ParseFloat(c.s(p), 64); return v }
