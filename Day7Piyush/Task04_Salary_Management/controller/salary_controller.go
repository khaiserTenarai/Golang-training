package controller

import (
	"fmt"
	"task04_salary_management/repository"
	"task04_salary_management/view"
)

type SalaryController struct {
	repo *repository.SalaryRepository
	view *view.SalaryView
}

func NewSalaryController(repo *repository.SalaryRepository, v *view.SalaryView) *SalaryController {
	return &SalaryController{repo: repo, view: v}
}

func (c *SalaryController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.AddEmployee()
		case 2:
			c.ListEmployees()
		case 3:
			c.UpdateSalary()
		case 4:
			c.ViewHistory()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *SalaryController) AddEmployee() {
	emp := c.view.GetEmployeeInput()
	id, err := c.repo.CreateEmployee(emp)
	if err != nil {
		c.view.ShowError("adding employee", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Employee created with ID: %d", id))
}

func (c *SalaryController) ListEmployees() {
	employees, err := c.repo.GetAll()
	if err != nil {
		c.view.ShowError("listing employees", err)
		return
	}
	c.view.ShowEmployees(employees)
}

func (c *SalaryController) UpdateSalary() {
	id := c.view.GetID("Employee")
	emp, err := c.repo.GetByID(id)
	if err != nil {
		c.view.ShowError("fetching employee", err)
		return
	}
	c.view.ShowCurrentSalary(emp)
	newSalary, reason := c.view.GetSalaryUpdateInput()
	if err := c.repo.UpdateSalary(id, newSalary, reason); err != nil {
		c.view.ShowError("updating salary (transaction rolled back)", err)
		return
	}
	c.view.ShowSuccess("Salary updated successfully via transaction.")
}

func (c *SalaryController) ViewHistory() {
	id := c.view.GetID("Employee")
	history, err := c.repo.GetSalaryHistory(id)
	if err != nil {
		c.view.ShowError("fetching salary history", err)
		return
	}
	c.view.ShowHistory(history)
}
