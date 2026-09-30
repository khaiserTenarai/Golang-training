package main

import "fmt"

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Department struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Floor int    `json:"floor"`
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

type Employee struct {
	ID         int        `json:"id"`
	Person     Person     `json:"person"`
	Salary     float64    `json:"salary"`
	HireDate   string     `json:"hire_date"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
	Active     bool       `json:"active"`
}

func (e Employee) FullName() string {
	return e.Person.FullName()
}

func (e Employee) AnnualSalary() float64 {
	return e.Salary * 12
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee[%d] %s - %s (%s) $%.2f/mo",
		e.ID, e.FullName(), e.Department.Name, e.Address.City, e.Salary)
}

func (e *Employee) GiveRaise(amount float64) {
	e.Salary += amount
}

func (e *Employee) Deactivate() {
	e.Active = false
}

func (e *Employee) Activate() {
	e.Active = true
}

func (e *Employee) Relocate(a Address) {
	e.Address = a
}
