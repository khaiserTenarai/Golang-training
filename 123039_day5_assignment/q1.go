package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Department string
	Salary     float64
	Phone      string
	City       string
}

func (e Employee) Display() {
	fmt.Println("Employee ID:", e.ID)
	fmt.Println("Employee Name:", e.Name)
	fmt.Println("Employee Email:", e.Email)
	fmt.Println("Employee Age:", e.Age)
	fmt.Println("Department:", e.Department)
	fmt.Println("Salary:", e.Salary)
	fmt.Println("Phone:", e.Phone)
	fmt.Println("City:", e.City)
}

func (e *Employee) IncreaseSalary(amount float64) {
	e.Salary = e.Salary + amount
}

func (e Employee) IsEligibleForBonus() bool {
	return e.Salary < 60000
}

type EmployeeOperations interface {
	Display()
	IncreaseSalary(amount float64)
	IsEligibleForBonus() bool
}

func main() {

	employee := Employee{
		ID:         101,
		Name:       "Swathi",
		Email:      "swathi@gmail.com",
		Age:        25,
		Department: "IT",
		Salary:     50000,
		Phone:      "9876543210",
		City:       "Bangalore",
	}

	fmt.Println("========================================")
	fmt.Println("       EMPLOYEE MANAGEMENT SYSTEM")
	fmt.Println("========================================")

	fmt.Println("\nEmployee Details:")
	employee.Display()

	fmt.Println("\nChecking Bonus Eligibility:")

	if employee.IsEligibleForBonus() {
		fmt.Println("Employee is eligible for bonus")
	} else {
		fmt.Println("Employee is not eligible for bonus")
	}

	fmt.Println("\nIncreasing Salary by 5000:")

	employee.IncreaseSalary(5000)

	fmt.Println("Updated Salary:", employee.Salary)

	fmt.Println("\nUsing Employee Interface:")

	var operation EmployeeOperations
	operation = &employee

	operation.Display()

	fmt.Println("\nInterface Salary Increase:")
	operation.IncreaseSalary(5000)

	fmt.Println("Salary after interface update:", employee.Salary)

	fmt.Println("\nChecking Bonus through Interface:")

	if operation.IsEligibleForBonus() {
		fmt.Println("Employee is eligible for bonus")
	} else {
		fmt.Println("Employee is not eligible for bonus")
	}
}