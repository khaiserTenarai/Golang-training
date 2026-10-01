package assignment_task10

import "errors"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func CalculateBonus(salary float64) (float64, error) {
	if salary < 0 {
		return 0, errors.New("salary cannot be negative")
	}
	return salary * 0.10, nil
}

func UpdateSalary(emp *Employee, increment float64) error {
	if increment < 0 {
		return errors.New("increment cannot be negative")
	}
	emp.Salary += increment
	return nil
}