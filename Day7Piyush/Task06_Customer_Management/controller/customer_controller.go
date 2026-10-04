package controller

import (
	"fmt"
	"task06_customer_management/repository"
	"task06_customer_management/view"
)

type CustomerController struct {
	repo *repository.CustomerRepository
	view *view.CustomerView
}

func NewCustomerController(repo *repository.CustomerRepository, v *view.CustomerView) *CustomerController {
	return &CustomerController{repo: repo, view: v}
}

func (c *CustomerController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.AddCustomer()
		case 2:
			c.ListCustomers()
		case 3:
			c.GetCustomer()
		case 4:
			c.UpdateCustomer()
		case 5:
			c.DeleteCustomer()
		case 6:
			c.SearchCustomers()
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *CustomerController) AddCustomer() {
	cust := c.view.GetCustomerInput()
	id, err := c.repo.Create(cust)
	if err != nil {
		c.view.ShowError("adding customer", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Customer created with ID: %d", id))
}

func (c *CustomerController) ListCustomers() {
	page, pageSize := c.view.GetPaginationInput()
	customers, total, err := c.repo.GetAll(page, pageSize)
	if err != nil {
		c.view.ShowError("listing customers", err)
		return
	}
	c.view.ShowCustomers(customers, total, page, pageSize)
}

func (c *CustomerController) GetCustomer() {
	id := c.view.GetID()
	cust, err := c.repo.GetByID(id)
	if err != nil {
		c.view.ShowError("fetching customer", err)
		return
	}
	c.view.ShowCustomer(cust)
}

func (c *CustomerController) UpdateCustomer() {
	id := c.view.GetID()
	existing, err := c.repo.GetByID(id)
	if err != nil {
		c.view.ShowError("fetching customer", err)
		return
	}
	updated := c.view.GetUpdateInput(existing)
	if err := c.repo.Update(updated); err != nil {
		c.view.ShowError("updating customer", err)
		return
	}
	c.view.ShowSuccess("Customer updated.")
}

func (c *CustomerController) DeleteCustomer() {
	id := c.view.GetID()
	if err := c.repo.Delete(id); err != nil {
		c.view.ShowError("deleting customer", err)
		return
	}
	c.view.ShowSuccess("Customer deleted.")
}

func (c *CustomerController) SearchCustomers() {
	keyword := c.view.GetSearchKeyword()
	customers, err := c.repo.Search(keyword)
	if err != nil {
		c.view.ShowError("searching customers", err)
		return
	}
	c.view.ShowSearchResults(customers)
}
