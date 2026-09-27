package main

import "fmt"

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Age        int     `json:"age"`
	Salary     float64 `json:"salary"`
	Phone      string  `json:"phone"`
	Position   string  `json:"position"`
	Experience int     `json:"experience"`
}

func main() {
	employee := Employee{
		ID:         101,
		Name:       "John",
		Email:      "john@gmail.com",
		Age:        28,
		Salary:     50000,
		Phone:      "9876543210",
		Position:   "Developer",
		Experience: 5,
	}

	fmt.Println(employee)
}