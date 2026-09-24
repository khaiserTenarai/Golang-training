package employee_test

import (
	"fmt"

	"godocdemo/employee"
)

// Example functions appear in the generated documentation and are
// verified by `go test` (the Output comment must match).
func ExampleStore_Add() {
	s := employee.NewStore()
	e := s.Add("Asha Rao", "Engineering", 85000)
	fmt.Println(e.ID, e.Name)
	// Output: 1 Asha Rao
}
