package controller

import (
	"fmt"
	"task10_order_inventory_transaction/repository"
	"task10_order_inventory_transaction/view"
)

type OrderController struct {
	repo *repository.OrderRepository
	view *view.OrderView
}

func NewOrderController(repo *repository.OrderRepository, v *view.OrderView) *OrderController {
	return &OrderController{repo: repo, view: v}
}

func (c *OrderController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.AddProduct()
		case 2:
			c.ListProducts()
		case 3:
			c.PlaceOrder()
		case 4:
			c.ListOrders()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *OrderController) AddProduct() {
	p := c.view.GetProductInput()
	id, err := c.repo.CreateProduct(p)
	if err != nil {
		c.view.ShowError("adding product", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Product added with ID: %d", id))
}

func (c *OrderController) ListProducts() {
	products, err := c.repo.GetAllProducts()
	if err != nil {
		c.view.ShowError("listing products", err)
		return
	}
	c.view.ShowProducts(products)
}

func (c *OrderController) PlaceOrder() {
	productID, quantity := c.view.GetOrderInput()
	orderID, err := c.repo.PlaceOrder(productID, quantity)
	if err != nil {
		c.view.ShowError("placing order", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Order placed with ID: %d (stock auto-reduced)", orderID))
}

func (c *OrderController) ListOrders() {
	orders, err := c.repo.GetAllOrders()
	if err != nil {
		c.view.ShowError("listing orders", err)
		return
	}
	c.view.ShowOrders(orders)
}
