package database

import (
	"sync"

	"day16-employee-graphql/model"
)

type DB struct {
	Mu        sync.RWMutex
	Employees map[int]model.Employee
	NextID    int
}

func New() *DB {
	return &DB{
		Employees: map[int]model.Employee{
			1: {ID: 1, Name: "Abirami", Email: "abirami@example.com", Age: 22, Salary: 45000, DepartmentID: 1},
			2: {ID: 2, Name: "Karthik", Email: "karthik@example.com", Age: 24, Salary: 52000, DepartmentID: 2},
		},
		NextID: 3,
	}
}
