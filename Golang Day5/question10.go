package main

import "fmt"

type Person struct {
	Name  string
	Email string
}

func (p Person) GetContactDetails() string {
	return p.Name + " <" + p.Email + ">"
}

type Employee struct {
	Person
	EmployeeID string
	Salary     int
}

type Contractor struct {
	Person
	HourlyRate int
	AgencyName string
}

func main() {
	emp := Employee{
		Person: Person{
			Name:  "Safa",
			Email: "safa@company.com",
		},
		EmployeeID: "EMP-042",
		Salary:     85000,
	}

	contractor := Contractor{
		Person: Person{
			Name:  "Riya",
			Email: "riya.r@agency.com",
		},
		HourlyRate: 75,
		AgencyName: "TechTalent",
	}

	fmt.Println(emp.Name)
	fmt.Println(emp.EmployeeID)
	fmt.Println(emp.GetContactDetails())

	fmt.Println(contractor.Name)
	fmt.Println(contractor.AgencyName)
	fmt.Println(contractor.GetContactDetails())
}