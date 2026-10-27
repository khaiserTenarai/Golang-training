package saga

import "fmt"

// Task 12: distributed transaction handling. Task 13: Saga workflow.
// This is a conceptual, in-memory example; it is not wired to the HTTP services or a database.
type EmployeeCreator interface {
	Create(name string) (int, error)
	CompensateDelete(employeeID int) error
}
type DepartmentAssigner interface {
	Assign(employeeID, departmentID int) error
}

type Result struct {
	EmployeeID int
	Completed  bool
}

// RegisterEmployee demonstrates orchestration and compensation.
// If assignment fails after employee creation, it compensates by deleting the employee.
func RegisterEmployee(ec EmployeeCreator, da DepartmentAssigner, name string, departmentID int) (Result, error) {
	id, err := ec.Create(name)
	if err != nil {
		return Result{}, fmt.Errorf("create employee: %w", err)
	}
	if err := da.Assign(id, departmentID); err != nil {
		if compErr := ec.CompensateDelete(id); compErr != nil {
			return Result{EmployeeID: id}, fmt.Errorf("assignment failed: %v; compensation also failed: %w", err, compErr)
		}
		return Result{EmployeeID: id}, fmt.Errorf("assignment failed; employee creation compensated: %w", err)
	}
	return Result{EmployeeID: id, Completed: true}, nil
}
