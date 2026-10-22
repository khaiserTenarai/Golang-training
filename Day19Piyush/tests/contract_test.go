// Task 14: Contract Tests — verify API contracts with mock servers.
// Run: go test -v ./tests/...

package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go_microservices/dto"
)

func mockDeptService() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/departments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.APIResponse{
			Success: true,
			Data:    []dto.DepartmentResponse{{ID: 1, Name: "Engineering", Location: "A", Budget: 500000}},
		})
	})
	mux.HandleFunc("/departments/1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.APIResponse{
			Success: true, Data: dto.DepartmentResponse{ID: 1, Name: "Engineering", Location: "A", Budget: 500000},
		})
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(dto.HealthResponse{Service: "department-service", Status: "healthy", Version: "1.0.0"})
	})
	return httptest.NewServer(mux)
}

func TestContract_DeptList(t *testing.T) {
	s := mockDeptService()
	defer s.Close()
	resp, _ := http.Get(s.URL + "/departments")
	var r dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&r)
	if !r.Success {
		t.Error("expected success")
	}
}

func TestContract_DeptByID(t *testing.T) {
	s := mockDeptService()
	defer s.Close()
	resp, _ := http.Get(s.URL + "/departments/1")
	var r dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&r)
	if !r.Success {
		t.Error("expected success")
	}
	b, _ := json.Marshal(r.Data)
	var d dto.DepartmentResponse
	json.Unmarshal(b, &d)
	if d.Name != "Engineering" {
		t.Errorf("expected Engineering, got %s", d.Name)
	}
}

func TestContract_Health(t *testing.T) {
	s := mockDeptService()
	defer s.Close()
	resp, _ := http.Get(s.URL + "/health")
	var h dto.HealthResponse
	json.NewDecoder(resp.Body).Decode(&h)
	if h.Status != "healthy" {
		t.Errorf("expected healthy, got %s", h.Status)
	}
}

func mockEmpService() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(dto.APIResponse{
				Success: true,
				Data:    []dto.EmployeeResponse{{ID: 1, Name: "Piyush", Email: "p@t.com", Salary: 60000}},
			})
		case "POST":
			var req dto.CreateEmployeeRequest
			json.NewDecoder(r.Body).Decode(&req)
			if req.Name == "" {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(dto.ErrorResponse{Error: "name required", Code: 400})
				return
			}
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(dto.APIResponse{Success: true, Data: dto.EmployeeResponse{ID: 99, Name: req.Name}})
		}
	})
	return httptest.NewServer(mux)
}

func TestContract_EmpList(t *testing.T) {
	s := mockEmpService()
	defer s.Close()
	resp, _ := http.Get(s.URL + "/employees")
	var r dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&r)
	if !r.Success {
		t.Error("expected success")
	}
}

func TestContract_EmpCreate(t *testing.T) {
	s := mockEmpService()
	defer s.Close()
	resp, _ := http.Post(s.URL+"/employees", "application/json",
		strings.NewReader(`{"name":"Test","email":"t@t.com","age":25,"department_id":1,"salary":50000}`))
	if resp.StatusCode != 201 {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
}

func TestContract_EmpCreateValidation(t *testing.T) {
	s := mockEmpService()
	defer s.Close()
	resp, _ := http.Post(s.URL+"/employees", "application/json", strings.NewReader(`{"name":""}`))
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestContract_CorrelationID(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := r.Header.Get("X-Correlation-ID")
		w.Header().Set("X-Correlation-ID", cid)
		json.NewEncoder(w).Encode(map[string]string{"cid": cid})
	}))
	defer s.Close()
	req, _ := http.NewRequest("GET", s.URL+"/test", nil)
	req.Header.Set("X-Correlation-ID", "test-cid-123")
	resp, _ := http.DefaultClient.Do(req)
	if resp.Header.Get("X-Correlation-ID") != "test-cid-123" {
		t.Error("correlation ID not propagated")
	}
}
