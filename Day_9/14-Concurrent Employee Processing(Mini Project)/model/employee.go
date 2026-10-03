package model
/*
	Employee represents one employee.

	Day 7:
	ID
	Name
	Age
	Email

	Day 9:
	Salary is added because the employee
	processing operation works with salary.
*/
type Employee struct {
	ID     int
	Name   string
	Age    int
	Email  string
	Salary float64
}