// Task 14: Test Gin handlers using httptest

package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"advanced_gin/models"

	"github.com/gin-gonic/gin"
)

func setupRouter() (*gin.Engine, *EmployeeHandler) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewEmployeeHandler()
	api := r.Group("/api")
	api.GET("/employees", h.GetAll)
	api.GET("/employees/:id", h.GetByID)
	api.POST("/employees", h.Create)
	api.PUT("/employees/:id", h.Update)
	api.DELETE("/employees/:id", h.Delete)
	return r, h
}

func TestGetAllEmployees(t *testing.T) {
	r, _ := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/employees", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp models.PaginatedResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if !resp.Success {
		t.Error("expected success=true")
	}
	if resp.Total != 5 {
		t.Errorf("expected 5 seeded employees, got %d", resp.Total)
	}
}

func TestGetEmployeeByID(t *testing.T) {
	r, _ := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/employees/1", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetEmployeeNotFound(t *testing.T) {
	r, _ := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/employees/999", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

func TestCreateEmployee(t *testing.T) {
	r, _ := setupRouter()
	body := `{"name":"Eve","email":"eve@test.com","age":26,"department":"Finance","salary":52000}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/employees", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}

func TestCreateEmployeeValidationFail(t *testing.T) {
	r, _ := setupRouter()
	body := `{"name":"A","email":"bad","age":10}` // fails validation
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/employees", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestUpdateEmployee(t *testing.T) {
	r, _ := setupRouter()
	body := `{"name":"Piyush Updated","email":"p@new.com","age":26,"department":"Backend","salary":70000}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/employees/1", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestDeleteEmployee(t *testing.T) {
	r, _ := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/employees/5", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestDeleteNotFound(t *testing.T) {
	r, _ := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/employees/999", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

func TestPagination(t *testing.T) {
	r, _ := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/employees?page=1&page_size=2", nil)
	r.ServeHTTP(w, req)

	var resp models.PaginatedResponse
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.PageSize != 2 {
		t.Errorf("expected page_size 2, got %d", resp.PageSize)
	}
	if resp.TotalPages != 3 {
		t.Errorf("expected 3 total pages for 5 items with size 2, got %d", resp.TotalPages)
	}
}
