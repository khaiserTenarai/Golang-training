package main

import (
	"sync"
	"testing"
)

func TestEmployeeService(t *testing.T) {
	service := NewEmployeeService()
	emp := Employee{ID: 1, Name: "Alice", Age: 30}

	err := service.AddEmployee(emp)
	if err != nil {
		t.Fatalf("failed to add employee: %v", err)
	}

	got, err := service.GetEmployee(1)
	if err != nil || got.Name != "Alice" {
		t.Errorf("failed to get employee")
	}
}

func TestEmployeeServiceConcurrentAccess(t *testing.T) {
	service := NewEmployeeService()
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_ = service.AddEmployee(Employee{ID: id, Name: "User", Age: 20})
			_, _ = service.GetEmployee(id)
		}(i)
	}

	wg.Wait()
}