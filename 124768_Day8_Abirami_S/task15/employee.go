package main

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func ValidateEmployee(employee Employee) bool {
	if employee.ID <= 0 {
		return false
	}

	if employee.Name == "" {
		return false
	}

	if employee.Salary < 0 {
		return false
	}

	return true
}
