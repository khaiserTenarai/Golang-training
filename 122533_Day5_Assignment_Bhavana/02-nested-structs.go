// 2. Create nested Address and Department structs.

package main

import "fmt"

type Address struct {
	Street  string
	City    string
	Pincode string
}

type Department struct {
	Name string
	Code string
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
	emp := Employee{
		ID:     1,
		Name:   "Anita",
		Email:  "anita@example.com",
		Age:    28,
		Salary: 45000,
		Phone:  "9876543210",
		Address: Address{
			Street:  "12 MG Road",
			City:    "Bangalore",
			Pincode: "560001",
		},
		Department: Department{
			Name: "Engineering",
			Code: "ENG",
		},
	}

	fmt.Println("Employee Name:", emp.Name)
	fmt.Println("City:", emp.Address.City)
	fmt.Println("Department:", emp.Department.Name)
	fmt.Printf("%+v\n", emp)
}
