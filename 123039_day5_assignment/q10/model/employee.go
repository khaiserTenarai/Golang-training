package model

import "fmt"

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode int    `json:"pincode"`
}

type Department struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Age        int        `json:"age"`
	Salary     float64    `json:"salary"`
	Phone      string     `json:"phone"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
}

func (e Employee) Display() {
	fmt.Println("Employee ID:", e.ID)
	fmt.Println("Employee Name:", e.Name)
	fmt.Println("Employee Email:", e.Email)
	fmt.Println("Employee Age:", e.Age)
	fmt.Println("Employee Salary:", e.Salary)
	fmt.Println("Employee Phone:", e.Phone)
	fmt.Println("City:", e.Address.City)
	fmt.Println("Department:", e.Department.Name)
}

func (e *Employee) IncreaseSalary(amount float64) {
	e.Salary = e.Salary + amount
}

func (e Employee) IsEligibleForBonus() bool {
	return e.Salary < 60000
}
