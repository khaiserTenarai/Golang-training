package main

import (
	"sync"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	service := NewEmployeeService()

	id := service.Add("Anita", 45000)

	emp, err := service.Get(id)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if emp.Name != "Anita" {
		t.Errorf("expected name Anita, got %s", emp.Name)
	}
	if emp.Salary != 45000 {
		t.Errorf("expected salary 45000, got %v", emp.Salary)
	}
}

func TestGetNotFound(t *testing.T) {
	service := NewEmployeeService()

	_, err := service.Get(99)
	if err == nil {
		t.Error("expected an error for a missing employee, got nil")
	}
}

func TestDelete(t *testing.T) {
	service := NewEmployeeService()
	id := service.Add("Ravi", 60000)

	if err := service.Delete(id); err != nil {
		t.Fatalf("expected no error deleting, got %v", err)
	}

	_, err := service.Get(id)
	if err == nil {
		t.Error("expected an error getting a deleted employee, got nil")
	}
}

func TestDeleteNotFound(t *testing.T) {
	service := NewEmployeeService()

	err := service.Delete(99)
	if err == nil {
		t.Error("expected an error deleting a missing employee, got nil")
	}
}

func TestCount(t *testing.T) {
	service := NewEmployeeService()
	service.Add("Anita", 45000)
	service.Add("Ravi", 60000)

	if got := service.Count(); got != 2 {
		t.Errorf("expected count 2, got %d", got)
	}
}

// TestConcurrentAdd calls Add() from many goroutines at the same time.
// Run this specifically with "go test -race" - if the mutex inside
// EmployeeService were ever removed, this test would trigger a data race
// warning (and likely an incorrect final count too).
func TestConcurrentAdd(t *testing.T) {
	service := NewEmployeeService()

	var wg sync.WaitGroup
	goroutines := 50

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			service.Add("Employee", float64(n)*1000)
		}(i)
	}

	wg.Wait()

	if got := service.Count(); got != goroutines {
		t.Errorf("expected %d employees, got %d", goroutines, got)
	}
}
