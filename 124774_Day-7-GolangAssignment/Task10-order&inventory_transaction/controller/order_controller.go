package controller

import (
	"fmt"

	"order-inventory/service"
	"order-inventory/view"
)

type OrderController struct {
	service service.OrderService
	view    *view.OrderView
}

func NewOrderController(
	service service.OrderService,
	view *view.OrderView,
) *OrderController {

	return &OrderController{
		service: service,
		view:    view,
	}
}

func (c *OrderController) Start() {

	for {

		c.view.ShowMenu()

		choice := c.view.ReadInt(
			"Enter your choice: ",
		)

		switch choice {

		case 1:
			c.CreateOrder()

		case 2:
			fmt.Println("Thank you. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *OrderController) CreateOrder() {

	customerName := c.view.ReadString(
		"Enter Customer Name: ",
	)

	productID := c.view.ReadInt(
		"Enter Product ID: ",
	)

	quantity := c.view.ReadInt(
		"Enter Quantity: ",
	)

	err := c.service.CreateOrder(
		customerName,
		productID,
		quantity,
	)

	if err != nil {

		c.view.ShowMessage(
			"Error: " + err.Error(),
		)

		return
	}

	c.view.ShowMessage(
		"Order created and stock updated successfully.",
	)
}
