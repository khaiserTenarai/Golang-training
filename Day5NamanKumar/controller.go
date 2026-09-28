package main

import (
	"encoding/json"
	"fmt"
)

type EmployeeController struct {
	service EmployeeService
}

func NewEmployeeController(s EmployeeService) *EmployeeController {
	return &EmployeeController{service: s}
}

func (c *EmployeeController) CreateEmployee(e Employee) {
	created, err := c.service.AddEmployee(e)
	if err != nil {
		fmt.Println("error creating employee:", err)
		return
	}
	fmt.Println("created:", created)
}

func (c *EmployeeController) ShowEmployee(id int) {
	e, err := c.service.GetEmployee(id)
	if err != nil {
		fmt.Println("error fetching employee:", err)
		return
	}
	fmt.Println(e)
}

func (c *EmployeeController) ShowEmployeeJSON(id int) {
	e, err := c.service.GetEmployee(id)
	if err != nil {
		fmt.Println("error fetching employee:", err)
		return
	}
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		fmt.Println("error marshaling employee:", err)
		return
	}
	fmt.Println(string(data))
}

func (c *EmployeeController) ListEmployees() {
	employees := c.service.ListEmployees()
	for _, e := range employees {
		fmt.Println(e)
	}
}

func (c *EmployeeController) GiveRaise(id int, amount float64) {
	e, err := c.service.RaiseSalary(id, amount)
	if err != nil {
		fmt.Println("error giving raise:", err)
		return
	}
	fmt.Println("updated:", e)
}

func (c *EmployeeController) DeleteEmployee(id int) {
	if err := c.service.RemoveEmployee(id); err != nil {
		fmt.Println("error deleting employee:", err)
		return
	}
	fmt.Println("deleted employee", id)
}
