package employee

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

var employees []Employee

var employeeMap = make(map[int]Employee)
