package employee

import "testing"

func TestAddEmployee(t *testing.T) {
	svc := NewService()

	err := svc.Add(Employee{ID: 1, Name: "Lakshmi", Position: "Developer", Salary: 75000})
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	err = svc.Add(Employee{ID: 1, Name: "Ayush", Position: "Manager", Salary: 90000})
	if err == nil {
		t.Error("expected error for duplicate ID, got nil")
	}

	err = svc.Add(Employee{ID: 2, Name: "", Position: "Intern", Salary: 30000})
	if err == nil {
		t.Error("expected error for empty name, got nil")
	}
}

func TestGetEmployee(t *testing.T) {
	svc := NewService()
	svc.Add(Employee{ID: 1, Name: "Lakshmi", Position: "Developer", Salary: 75000})

	emp, err := svc.Get(1)
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
	if emp.Name != "Alice" {
		t.Errorf("got %s, want Lakshmi", emp.Name)
	}

	_, err = svc.Get(99)
	if err == nil {
		t.Error("expected error for missing employee, got nil")
	}
}

//go test -v ./...
//go fmt ./...
//go vet ./...
//go test -race ./...
//go test -v ./...
//go test -coverprofile coverage.out .
