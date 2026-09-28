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

	fmt.Println("Employee ID:", employee.ID)
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Employee Email:", employee.Email)
	fmt.Println("Employee Age:", employee.Age)
	fmt.Println("Employee Salary:", employee.Salary)
	fmt.Println("Employee Phone:", employee.Phone)

	fmt.Println("\nAddress:")
	fmt.Println("Street:", employee.Address.Street)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("State:", employee.Address.State)
	fmt.Println("Pincode:", employee.Address.Pincode)

	fmt.Println("\nDepartment:")
	fmt.Println("Department Name:", employee.Department.Name)
	fmt.Println("Manager:", employee.Department.Manager)
}
