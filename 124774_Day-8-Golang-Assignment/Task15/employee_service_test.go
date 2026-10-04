package employee

import (
	"sync"
	"testing"
)

func TestAddEmployee(t *testing.T) {
	service := NewEmployeeService()

	employee := Employee{
		ID:         1,
		Name:       "John",
		Email:      "john@example.com",
		Department: "IT",
		Salary:     50000,
	}

	err := service.AddEmployee(employee)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result, err := service.GetEmployee(1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Name != "John" {
		t.Errorf("expected John, got %s", result.Name)
	}
}

func TestAddDuplicateEmployee(t *testing.T) {
	service := NewEmployeeService()

	employee := Employee{
		ID:    1,
		Name:  "John",
		Email: "john@example.com",
	}

	_ = service.AddEmployee(employee)

	err := service.AddEmployee(employee)

	if err == nil {
		t.Error("expected duplicate employee error")
	}
}

func TestAddInvalidEmployee(t *testing.T) {
	service := NewEmployeeService()

	employee := Employee{
		ID:    0,
		Name:  "John",
		Email: "john@example.com",
	}

	err := service.AddEmployee(employee)

	if err == nil {
		t.Error("expected invalid employee ID error")
	}
}

func TestGetEmployeeNotFound(t *testing.T) {
	service := NewEmployeeService()

	_, err := service.GetEmployee(100)

	if err == nil {
		t.Error("expected employee not found error")
	}
}

func TestGetAllEmployees(t *testing.T) {
	service := NewEmployeeService()

	employees := []Employee{
		{
			ID:    1,
			Name:  "John",
			Email: "john@example.com",
		},
		{
			ID:    2,
			Name:  "Jane",
			Email: "jane@example.com",
		},
	}

	for _, employee := range employees {
		if err := service.AddEmployee(employee); err != nil {
			t.Fatal(err)
		}
	}

	result := service.GetAllEmployees()

	if len(result) != 2 {
		t.Errorf("expected 2 employees, got %d", len(result))
	}
}

func TestUpdateEmployee(t *testing.T) {
	service := NewEmployeeService()

	employee := Employee{
		ID:    1,
		Name:  "John",
		Email: "john@example.com",
	}

	_ = service.AddEmployee(employee)

	employee.Name = "John Updated"

	err := service.UpdateEmployee(employee)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result, _ := service.GetEmployee(1)

	if result.Name != "John Updated" {
		t.Errorf("expected updated name, got %s", result.Name)
	}
}

func TestUpdateEmployeeNotFound(t *testing.T) {
	service := NewEmployeeService()

	employee := Employee{
		ID:    100,
		Name:  "Unknown",
		Email: "unknown@example.com",
	}

	err := service.UpdateEmployee(employee)

	if err == nil {
		t.Error("expected employee not found error")
	}
}

func TestDeleteEmployee(t *testing.T) {
	service := NewEmployeeService()

	employee := Employee{
		ID:    1,
		Name:  "John",
		Email: "john@example.com",
	}

	_ = service.AddEmployee(employee)

	err := service.DeleteEmployee(1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = service.GetEmployee(1)

	if err == nil {
		t.Error("expected employee not found error")
	}
}

func TestDeleteEmployeeNotFound(t *testing.T) {
	service := NewEmployeeService()

	err := service.DeleteEmployee(100)

	if err == nil {
		t.Error("expected employee not found error")
	}
}

func TestRaceCondition(t *testing.T) {
	service := NewEmployeeService()

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			employee := Employee{
				ID:    id,
				Name:  "Employee",
				Email: "employee@example.com",
			}

			_ = service.AddEmployee(employee)
		}(i)
	}

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			_, _ = service.GetEmployee(id)
		}(i)
	}

	wg.Wait()

	employees := service.GetAllEmployees()

	if len(employees) != 100 {
		t.Errorf("expected 100 employees, got %d", len(employees))
	}
}
