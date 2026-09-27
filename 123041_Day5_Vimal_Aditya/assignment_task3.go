package main

import (
	"fmt"
)

type Address1 struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}

type Department1 struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

type Employee3 struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Salary     float64    `json:"salary"`
	Address    Address1    `json:"address"`
	Department Department1 `json:"department"`
}

func main() {

	fmt.Println("\n**********************")
	fmt.Println("3. Add JSON struct tags.")
	fmt.Println("************************")

	emp := Employee3{
		ID:     101,
		Name:   "Vimal Aditya",
		Salary: 500000,
		Address: Address1{
			City:    "Bangalore",
			State:   "Karnataka",
			Pincode: "560016",
		},
		Department: Department1{
			Name:    "Engineering",
			Manager: "Alex",
		},
	}

	fmt.Println("ID        :", emp.ID)
	fmt.Println("Name      :", emp.Name)
	fmt.Println("Salary    :", emp.Salary)
	fmt.Println("City      :", emp.Address.City)
	fmt.Println("Department:", emp.Department.Name)
	fmt.Println("Manager   :", emp.Department.Manager)
}