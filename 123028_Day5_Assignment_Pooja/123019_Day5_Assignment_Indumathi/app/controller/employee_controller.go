package controller

import (
	"fmt"

	"app/model"
	"app/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) Create(emp *model.Employee) {
	err := c.service.RegisterEmployee(emp)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
		return
	}
	fmt.Printf("[Success] Created employee: %s %s (ID: %d)\n", emp.FirstName, emp.LastName, emp.ID)
}

func (c *EmployeeController) Show(id int) {
	emp, err := c.service.GetEmployee(id)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
		return
	}
	fmt.Printf("[Info] ID: %d | Name: %s %s | Salary: $%.2f | Active: %t\n",
		emp.ID, emp.FirstName, emp.LastName, emp.Salary, emp.IsActive)
}

func (c *EmployeeController) ProcessRaise(id int, percent float64) {
	err := c.service.GiveRaise(id, percent)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
		return
	}
	fmt.Printf("[Success] Granted %.1f%% raise to ID %d\n", percent, id)
}

func (c *EmployeeController) Deactivate(id int) {
	err := c.service.DeactivateEmployee(id)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
		return
	}
	fmt.Printf("[Success] Deactivated ID %d\n", id)
}