package model

type Employee struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Salary float64
}

func (e Employee) Display() {
	println("ID:", e.ID)
	println("Name:", e.Name)
	println("Email:", e.Email)
	println("Age:", e.Age)
	println("Salary:", e.Salary)
}