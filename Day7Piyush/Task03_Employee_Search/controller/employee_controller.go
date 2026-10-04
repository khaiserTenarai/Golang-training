package controller

import (
	"fmt"
	"task03_employee_search/repository"
	"task03_employee_search/view"
)

type EmployeeController struct {
	repo *repository.EmployeeRepository
	view *view.EmployeeView
}

func NewEmployeeController(repo *repository.EmployeeRepository, v *view.EmployeeView) *EmployeeController {
	return &EmployeeController{repo: repo, view: v}
}

func (c *EmployeeController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.AddEmployee()
		case 2:
			c.SearchEmployees()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *EmployeeController) AddEmployee() {
	emp := c.view.GetEmployeeInput()
	id, err := c.repo.Insert(emp)
	if err != nil {
		c.view.ShowError("adding employee", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Employee added with ID: %d", id))
}

func (c *EmployeeController) SearchEmployees() {
	params := c.view.GetSearchParams()
	employees, total, err := c.repo.Search(params)
	if err != nil {
		c.view.ShowError("searching employees", err)
		return
	}
	page := params.Page
	if page < 1 {
		page = 1
	}
	pageSize := params.PageSize
	if pageSize < 1 {
		pageSize = 5
	}
	c.view.ShowEmployees(employees, total, page, pageSize)
}
