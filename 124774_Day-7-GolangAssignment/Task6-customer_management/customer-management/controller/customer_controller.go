package controller

import (
	"fmt"

	"customer-management/service"
	"customer-management/view"
)

type CustomerController struct {
	service service.CustomerService
	view    *view.CustomerView
}

// ==================================================
// CONSTRUCTOR
// ==================================================

func NewCustomerController(
	service service.CustomerService,
	view *view.CustomerView,
) *CustomerController {

	return &CustomerController{
		service: service,
		view:    view,
	}
}

// ==================================================
// START
// ==================================================

func (c *CustomerController) Start() {

	for {

		c.view.ShowMenu()

		choice := c.view.ReadInt(
			"Enter your choice: ",
		)

		switch choice {

		case 1:
			c.AddCustomer()

		case 2:
			c.FindCustomerByID()

		case 3:
			c.FindAllCustomers()

		case 4:
			c.UpdateCustomer()

		case 5:
			c.DeleteCustomer()

		case 6:
			c.SearchCustomers()

		case 7:
			c.FindCustomersByPage()

		case 8:
			fmt.Println(
				"Thank you. Goodbye!",
			)

			return

		default:
			fmt.Println(
				"Invalid choice.",
			)
		}
	}
}

// ==================================================
// CREATE
// ==================================================

func (c *CustomerController) AddCustomer() {

	customer :=
		c.view.ReadCustomer()

	err :=
		c.service.AddCustomer(customer)

	if err != nil {

		c.view.ShowMessage(
			"Error: " + err.Error(),
		)

		return
	}

	c.view.ShowMessage(
		"Customer added successfully.",
	)
}

// ==================================================
// FIND BY ID
// ==================================================

func (c *CustomerController) FindCustomerByID() {

	id := c.view.ReadInt(
		"Enter Customer ID: ",
	)

	customer, err :=
		c.service.FindCustomerByID(id)

	if err != nil {

		c.view.ShowMessage(
			"Error: " + err.Error(),
		)

		return
	}

	c.view.ShowCustomer(customer)
}

// ==================================================
// FIND ALL
// ==================================================

func (c *CustomerController) FindAllCustomers() {

	customers :=
		c.service.FindAllCustomers()

	c.view.ShowCustomers(customers)
}

// ==================================================
// UPDATE
// ==================================================

func (c *CustomerController) UpdateCustomer() {

	customer :=
		c.view.ReadCustomerForUpdate()

	err :=
		c.service.UpdateCustomer(customer)

	if err != nil {

		c.view.ShowMessage(
			"Error: " + err.Error(),
		)

		return
	}

	c.view.ShowMessage(
		"Customer updated successfully.",
	)
}

// ==================================================
// DELETE
// ==================================================

func (c *CustomerController) DeleteCustomer() {

	id := c.view.ReadInt(
		"Enter Customer ID: ",
	)

	err :=
		c.service.DeleteCustomer(id)

	if err != nil {

		c.view.ShowMessage(
			"Error: " + err.Error(),
		)

		return
	}

	c.view.ShowMessage(
		"Customer deleted successfully.",
	)
}

// ==================================================
// SEARCH
// ==================================================

func (c *CustomerController) SearchCustomers() {

	keyword := c.view.ReadString(
		"Enter search keyword: ",
	)

	customers :=
		c.service.SearchCustomers(keyword)

	c.view.ShowCustomers(customers)
}

// ==================================================
// PAGINATION
// ==================================================

func (c *CustomerController) FindCustomersByPage() {

	page := c.view.ReadInt(
		"Enter page number: ",
	)

	pageSize := c.view.ReadInt(
		"Enter page size: ",
	)

	customers :=
		c.service.FindCustomersByPage(
			page,
			pageSize,
		)

	c.view.ShowCustomers(customers)
}
