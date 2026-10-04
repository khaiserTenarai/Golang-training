package controller

import (
	"fmt"

	"employee-managementPostgre/model"
	"employee-managementPostgre/service"
)

type EmployeeController struct {
	Service service.EmployeeService
}

func (c *EmployeeController) AddEmployee(employee *model.Employee) error {

	err := c.Service.AddEmployee(employee)

	if err != nil {
		return err
	}

	return nil
}

func (c *EmployeeController) GetEmployees() ([]model.Employee, error) {

	employees, err := c.Service.GetEmployees()

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (c *EmployeeController) GetEmployee(id int) (model.Employee, error) {

	employee, err := c.Service.GetEmployee(id)

	if err != nil {
		return model.Employee{}, err
	}

	return employee, nil
}

func (c *EmployeeController) UpdateEmployee(employee *model.Employee) error {

	err := c.Service.UpdateEmployee(employee)

	if err != nil {
		return err
	}

	return nil
}

func (c *EmployeeController) DeleteEmployee(id int) error {

	err := c.Service.DeleteEmployee(id)

	if err != nil {
		return err
	}

	return nil
}

func DisplayEmployee(employee model.Employee) {

	fmt.Println("--------------------------------")
	fmt.Println("Employee ID   :", employee.ID)
	fmt.Println("Name          :", employee.Name)
	fmt.Println("Email         :", employee.Email)
	fmt.Println("Age           :", employee.Age)
	fmt.Println("Salary        :", employee.Salary)
	fmt.Println("Department ID :", employee.DepartmentID)
	fmt.Println("--------------------------------")
}