package main

import (
	"log/slog"
	"os"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func main() {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)

	employee := Employee{
		ID:         101,
		Name:       "ganesh",
		Department: "IT",
		Salary:     25000,
	}

	logger.Info("employee created",
		"employee_id", employee.ID,
		"name", employee.Name,
		"department", employee.Department,
		"salary", employee.Salary,
	)
}
