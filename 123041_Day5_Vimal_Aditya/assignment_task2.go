package main

import "fmt"

type Address struct {
	City    string
	State   string
	Pincode string
}

type Department struct {
	Name    string
	Manager string
}

type Employee2 struct {
	ID         int
	Name       string
	Salary     float64
	Address    Address
	Department Department
}

func main() {

	fmt.Println("\n**********************************************")
	fmt.Println("2. Create nested Address and Department structs.")
	fmt.Println("************************************************")

	emp := Employee2{
		ID:     101,
		Name:   "Vimal Aditya",
		Salary: 500000,
		Address: Address{
			City:    "Bangalore",
			State:   "Karnataka",
			Pincode: "560016",
		},
		Department: Department{
			Name:    "Engineering",
			Manager: "Alex",
		},
	}

	fmt.Println("ID        :", emp.ID)
	fmt.Println("Name      :", emp.Name)
	fmt.Println("City      :", emp.Address.City)
	fmt.Println("Department:", emp.Department.Name)
	fmt.Println("Manager   :", emp.Department.Manager)
}