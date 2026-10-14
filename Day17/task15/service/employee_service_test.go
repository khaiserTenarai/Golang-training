package service_test

import (
	"testing"

	"employee-platform/models"
	"employee-platform/repository"
	"employee-platform/service"
)

func TestEmployeeService_CreateAndGet(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewEmployeeService(repo)

	emp := models.Employee{
		ID:           "test-1",
		Name:         "Tester",
		Role:         "QA",
		DepartmentID: "dept-1",
		Salary:       80000,
	}

	err := svc.CreateEmployee(&emp)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	retrieved, err := svc.GetEmployeeByID("test-1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if retrieved.Name != "Tester" {
		t.Errorf("expected Tester, got %s", retrieved.Name)
	}
}
