package model


type Employee struct {
	ID           int
	Name         string
	Email        string
	Age          int
	Salary       float64
	DepartmentID int
}

type Department struct {
	ID   int
	Name string
}
