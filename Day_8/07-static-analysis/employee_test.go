package main

import "testing"

func TestEmployee(t *testing.T) {
	employee := Employee{
		ID:         101,
		Name:       "Alice",
		Department: "Engineering",
		Salary:     75000,
	}

	if employee.ID != 101 {
		t.Errorf("expected ID 101, got %d", employee.ID)
	}

	if employee.Name != "Alice" {
		t.Errorf("expected Alice, got %s", employee.Name)
	}

	if employee.Department != "Engineering" {
		t.Errorf("expected Engineering, got %s", employee.Department)
	}
}

/*
Static Analysis Steps:
1. gofmt -w .       -> Formats the code
2. go vet ./...     -> Performs static analysis; no issues found
3. go test ./...    -> Tests passed
4. go test -cover ./... -> Coverage checked
5. go run .         -> Runs the program

Output:
ID: 101
Name: Alice
Department: Engineering
Salary: 75000.00
*/
