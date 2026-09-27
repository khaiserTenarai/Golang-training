package dao

import "emp_simple/model"

var employees []model.Employee

func AddEmp(employee *model.Employee) {
	employees = append(employees, *employee)
}

func GetEmployee() []model.Employee{
	return employees
}

func DeleteEmployee(id int) bool{
	for i, employee := range employees{
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return true
		}
	}
	return false
}



