package controller

import (
	"fmt"
	"task09_ecommerce_order_system/repository"
	"task09_ecommerce_order_system/view"
)

type OrderController struct {
	custRepo    *repository.CustomerRepository
	prodRepo    *repository.ProductRepository
	orderRepo   *repository.OrderRepository
	view        *view.OrderView
}

func NewOrderController(cr *repository.CustomerRepository, pr *repository.ProductRepository, or *repository.OrderRepository, v *view.OrderView) *OrderController {
	return &OrderController{custRepo: cr, prodRepo: pr, orderRepo: or, view: v}
}

func (c *OrderController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.AddCustomer()
		case 2:
			c.AddProduct()
		case 3:
			c.ListCustomers()
		case 4:
			c.ListProducts()
		case 5:
			c.CreateOrder()
		case 6:
			c.ViewOrderDetails()
		case 7:
			c.ListOrders()
		case 8:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *OrderController) AddCustomer() {
	cust := c.view.GetCustomerInput()
	id, err := c.custRepo.Create(cust)
	if err != nil {
		c.view.ShowError("adding customer", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Customer created with ID: %d", id))
}

func (c *OrderController) AddProduct() {
	prod := c.view.GetProductInput()
	id, err := c.prodRepo.Create(prod)
	if err != nil {
		c.view.ShowError("adding product", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Product created with ID: %d", id))
}

func (c *OrderController) ListCustomers() {
	customers, err := c.custRepo.GetAll()
	if err != nil {
		c.view.ShowError("listing customers", err)
		return
	}
	c.view.ShowCustomers(customers)
}

func (c *OrderController) ListProducts() {
	products, err := c.prodRepo.GetAll()
	if err != nil {
		c.view.ShowError("listing products", err)
		return
	}
	c.view.ShowProducts(products)
}

func (c *OrderController) CreateOrder() {
	custID, items := c.view.GetOrderItems()
	orderID, err := c.orderRepo.CreateOrder(custID, items)
	if err != nil {
		c.view.ShowError("creating order", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Order created with ID: %d", orderID))
}

func (c *OrderController) ViewOrderDetails() {
	id := c.view.GetID("Order")
	order, err := c.orderRepo.GetOrderDetails(id)
	if err != nil {
		c.view.ShowError("fetching order details", err)
		return
	}
	c.view.ShowOrderDetails(order)
}

func (c *OrderController) ListOrders() {
	orders, err := c.orderRepo.GetAllOrders()
	if err != nil {
		c.view.ShowError("listing orders", err)
		return
	}
	c.view.ShowOrders(orders)
}
