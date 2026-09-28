package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	Pincode int
}

type Department struct {
	Name    string
	Manager string
}

type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Salary     float64
	Phone      string
	Address    Address
	Department Department
}

func (e Employee) Display() {
	fmt.Println("Employee ID:", e.ID)
	fmt.Println("Employee Name:", e.Name)
	fmt.Println("Employee Email:", e.Email)
	fmt.Println("Employee Age:", e.Age)
	fmt.Println("Employee Salary:", e.Salary)
	fmt.Println("Employee Phone:", e.Phone)
}

func (e *Employee) IncreaseSalary(amount float64) {
	e.Salary = e.Salary + amount
}

func (e Employee) IsEligibleForBonus() bool {
	return e.Salary < 60000
}

func (e Employee) DisplayAddress() {
	fmt.Println("Street:", e.Address.Street)
	fmt.Println("City:", e.Address.City)
	fmt.Println("State:", e.Address.State)
	fmt.Println("Pincode:", e.Address.Pincode)
}

func (e Employee) DisplayDepartment() {
	fmt.Println("Department:", e.Department.Name)
	fmt.Println("Manager:", e.Department.Manager)
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Swathi",
		Email:  "swathi@gmail.com",
		Age:    25,
		Salary: 50000,
		Phone:  "9876543210",

		Address: Address{
			Street:  "MG Road",
			City:    "Bangalore",
			State:   "Karnataka",
			Pincode: 560001,
		},

		Department: Department{
			Name:    "IT",
			Manager: "Rahul",
		},
	}

	fmt.Println("===== EMPLOYEE DETAILS =====")
	employee.Display()

	fmt.Println("\n===== ADDRESS =====")
	employee.DisplayAddress()

	fmt.Println("\n===== DEPARTMENT =====")
	employee.DisplayDepartment()

	fmt.Println("\n===== SALARY =====")
	fmt.Println("Before:", employee.Salary)

	employee.IncreaseSalary(5000)

	fmt.Println("After:", employee.Salary)

	fmt.Println("\n===== BONUS =====")

	if employee.IsEligibleForBonus() {
		fmt.Println("Employee is eligible for bonus")
	} else {
		fmt.Println("Employee is not eligible for bonus")
	}
}