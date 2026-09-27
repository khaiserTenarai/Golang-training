package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	Pincode string
}

type Department struct {
	ID   int
	Name string
}

type Employee struct {
	ID          int
	Name        string
	Email       string
	Age         int
	Salary      float64
	Phone       string
	Position    string
	JoiningDate string
	Address     Address
	Department  Department
}

func main() {

	employee := Employee{
		ID:          101,
		Name:        "Muneera",
		Email:       "muneera@gmail.com",
		Age:         21,
		Salary:      20000,
		Phone:       "9876543210",
		Position:    "Software Developer",
		JoiningDate: "16-09-2026",

		Address: Address{
			City:    "Bangalore",
			State:   "Karnataka",
			Pincode: "560037",
		},

		Department: Department{
			ID:   10,
			Name: "IT",
		},
	}

	fmt.Println("------Employee Details------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("Phone:", employee.Phone)
	fmt.Println("Position:", employee.Position)
	fmt.Println("Joining Date:", employee.JoiningDate)

	fmt.Println("\n-----Address------")
	fmt.Println("Street:", employee.Address.Street)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("State:", employee.Address.State)
	fmt.Println("Pincode:", employee.Address.Pincode)

	fmt.Println("\n-----Department-----")
	fmt.Println("Department ID:", employee.Department.ID)
	fmt.Println("Department Name:", employee.Department.Name)
}
