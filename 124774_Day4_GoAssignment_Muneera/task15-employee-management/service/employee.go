package service

import (
	"fmt"
	"task15-employee-management/model"
	"task15-employee-management/utility"
)

var employees []model.Employee

func AddEmployee(employee model.Employee) error {
	err := utility.ValidateEmployeee(
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
	)
	if err != nil {
		return err
	}
	for i := 0; i < len(employees); i++ {
		if employees[i].Id == employee.Id {
			return utility.ErrDuplicateEmployee
		}
	}
	employee.Name = utility.CleanString(employee.Name)
	employee.Email = utility.CleanString(employee.Email)

	employees = append(employees, employee)

	fmt.Println("Employee Added successfully")
	return nil
}

func GetEmployee(id int) (model.Employee, error) {
	for i := 0; i < len(employees); i++ {
		if employees[i].Id == id {
			return employees[i], nil
		}
	}
	return model.Employee{}, utility.ErrEmployeeNotFound
}

func DeleteEmployee(id int) error {
	for i := 0; i < len(employees); i++ {
		if employees[i].Id == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted successfully")
			return nil
		}
	}
	return utility.ErrEmployeeNotFound

}
