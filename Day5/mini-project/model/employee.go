package model

import "fmt"

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Address struct {
	City    string `json:"city"`
	Pincode string `json:"pincode"`
}


type Department struct {
	Name string `json:"name"`
}

type Employee struct {
	Person     `json:"person"` // composition, not inheritance
	ID         int        `json:"id"`
	Salary     float64    `json:"salary"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
	IsActive   bool       `json:"is_active"`
}

// ---------- using Pointer receiver methods here ----------

func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * percent / 100
}

func (e *Employee) Deactivate() {
	e.IsActive = false
}

func (e *Employee) UpdateAddress(a Address) {
	e.Address = a
}

func (e Employee) String() string {
	status := "Active"
	if !e.IsActive {
		status = "Inactive"
	}
	return fmt.Sprintf("[%d] %s (%s) - %s dept, salary %.2f, %s, %s %s",
		e.ID, e.Name, e.Email, e.Department.Name, e.Salary, status, e.Address.City, e.Address.Pincode)
}

func (e Employee) IsSenior() bool {
	return e.Age >= 45
}

func (e Employee) AnnualSalary() float64 {
	return e.Salary * 12
}
