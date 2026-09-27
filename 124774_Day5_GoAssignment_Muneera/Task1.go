package main

import "fmt"

type Employee struct {
	ID          int
	Name        string
	Email       string
	Age         int
	Salary      float64
	Phone       string
	Position    string
	JoiningDate string
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
	}

	fmt.Println("Employee Details")
	fmt.Println("-------------------------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("Phone:", employee.Phone)
	fmt.Println("Position:", employee.Position)
	fmt.Println("Joining Date:", employee.JoiningDate)
}
