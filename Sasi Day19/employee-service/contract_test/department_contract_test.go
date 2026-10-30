package contract_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"employee-service/client"
	"employee-service/dto"
)

func TestDepartmentClientContract(t *testing.T) {
	expectedDept := dto.DepartmentDTO{
		ID:        1,
		Name:      "Engineering",
		Code:      "ENG-01",
		Headcount: 50,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/departments/1" {
			t.Errorf("Expected path /api/v1/departments/1, got %s", r.URL.Path)
		}
		if r.Header.Get("X-Correlation-ID") == "" {
			t.Errorf("Expected X-Correlation-ID header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedDept)
	}))
	defer server.Close()

	deptClient := client.NewDepartmentClient(server.URL)
	dept, err := deptClient.GetDepartmentByID(context.Background(), 1, "test-corr-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if dept.Name != expectedDept.Name {
		t.Errorf("Expected name %s, got %s", expectedDept.Name, dept.Name)
	}
}
