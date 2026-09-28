package main

import "fmt"

type Address struct {
	City    string
	State   string
	Pincode string
}
type Department struct {
	ID   int
	Name string
}
type Employee2 struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Salary     float64
	Department Department
	Address    Address
}

func main() {
	employee := Employee2{
		ID:     101,
		Name:   "Abirami",
		Email:  "abi@gmail.com",
		Age:    22,
		Salary: 90000,
		Department: Department{
			ID:   1,
			Name: "IT",
		},
		Address: Address{
			City:    "Banglore",
			State:   "Karnataka",
			Pincode: "506100",
		},
	}
	fmt.Println("Employee Details")
	fmt.Println("ID: ", employee.ID)
	fmt.Println("Name: ", employee.Name)
	fmt.Println("Email: ", employee.Email)
	fmt.Println("Age: ", employee.Age)
	fmt.Println("Salary: ", employee.Salary)

	fmt.Println("-------------------------------------")
	fmt.Println("Address")
	fmt.Println("City: ", employee.Address.City)
	fmt.Println("State: ", employee.Address.State)
	fmt.Println("Pincode: ", employee.Address.Pincode)

	fmt.Println("-------------------------------------")
	fmt.Println("Department")
	fmt.Println("ID: ", employee.Department.ID)
	fmt.Println("Name: ", employee.Department.Name)
}
