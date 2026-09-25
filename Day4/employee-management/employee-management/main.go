package main

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func doubleSalaryPointer(emp *model.Employee) {
	emp.Salary = emp.Salary * 2
}

func doubleSalaryValue(emp model.Employee) {
	emp.Salary = emp.Salary * 2
}

func main() {
	svc := service.NewEmployeeService()
	generateID := service.NewIDGenerator(1)

	employees := []model.Employee{
		{Name: "Alice Johnson", Email: "alice@example.com", Age: 28, Salary: 55000},
		{Name: "Bob Smith", Email: "bob@example.com", Age: 34, Salary: 62000},
		{Name: "Carol Lee", Email: "carol@example.com", Age: 41, Salary: 48000},
	}

	for i := range employees {
		employees[i].ID = generateID()
		if err := svc.AddEmployee(&employees[i]); err != nil {
			fmt.Println("Error adding employee:", err)
			continue
		}
		fmt.Printf("Added employee: %+v\n", employees[i])
	}

	dupErr := svc.AddEmployee(&employees[0])
	if dupErr != nil {
		fmt.Println("Duplicate check:", dupErr)
		if errors.Is(dupErr, service.ErrDuplicateEmployee) {
			fmt.Println("Confirmed: duplicate employee error")
		}
	}

	_, notFoundErr := svc.GetEmployee(999)
	if notFoundErr != nil {
		fmt.Println("Lookup failure:", notFoundErr)
		if errors.Is(notFoundErr, service.ErrEmployeeNotFound) {
			fmt.Println("Confirmed: employee not found error")
		}
	}

	invalidEmp := &model.Employee{ID: 100, Name: "A", Email: "bademail", Age: 10, Salary: -5}
	validationErr := svc.AddEmployee(invalidEmp)
	if validationErr != nil {
		fmt.Println("Validation failure:", validationErr)
		var valErr *utility.ValidationError
		if errors.As(validationErr, &valErr) {
			fmt.Printf("Field: %s, Message: %s\n", valErr.Field, valErr.Message)
		}
	}

	count, total := svc.EmployeeStats()
	fmt.Printf("Employee count: %d, Total salary: %.2f\n", count, total)

	combinedSalary := service.CalculateTotalSalary(55000, 62000, 48000, 30000)
	fmt.Println("Combined salary calculation:", combinedSalary)

	filterHighEarners := func(emps []*model.Employee, threshold float64) []*model.Employee {
		result := make([]*model.Employee, 0)
		for _, emp := range emps {
			if emp.Salary >= threshold {
				result = append(result, emp)
			}
		}
		return result
	}

	highEarners := filterHighEarners(svc.AllEmployees(), 50000)
	fmt.Println("High earners:")
	for _, emp := range highEarners {
		fmt.Printf("  %s - %.2f\n", emp.Name, emp.Salary)
	}

	fmt.Println("Factorial of 5:", factorial(5))

	fmt.Println("Fibonacci recursive(10):", fibonacciRecursive(10))
	fmt.Println("Fibonacci iterative(10):", fibonacciIterative(10))

	sample := model.Employee{Name: "Dan", Email: "dan@example.com", Age: 30, Salary: 40000}
	fmt.Println("Before pointer modification:", sample.Salary)
	doubleSalaryPointer(&sample)
	fmt.Println("After pointer modification:", sample.Salary)

	sample2 := model.Employee{Name: "Eve", Email: "eve@example.com", Age: 29, Salary: 40000}
	fmt.Println("Before value modification:", sample2.Salary)
	doubleSalaryValue(sample2)
	fmt.Println("After value modification (unchanged):", sample2.Salary)

	delErr := svc.DeleteEmployee(employees[1].ID)
	if delErr != nil {
		fmt.Println("Delete error:", delErr)
	} else {
		fmt.Println("Deleted employee ID:", employees[1].ID)
	}

	_, getErr := svc.GetEmployee(employees[1].ID)
	fmt.Println("Get after delete:", getErr)
}
