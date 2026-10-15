package controller

import (
	"bytes"
	"context"
	"employee-management/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	employees []model.Employee
}

func (m *mockService) CreateEmployee(ctx context.Context, emp *model.Employee) error {
	emp.ID = 1
	m.employees = append(m.employees, *emp)
	return nil
}

func (m *mockService) GetEmployees(ctx context.Context, page, limit int) ([]model.Employee, int, error) {
	return m.employees, len(m.employees), nil
}

func (m *mockService) GetEmployeeByID(ctx context.Context, id int64) (*model.Employee, error) {
	for _, e := range m.employees {
		if e.ID == id {
			return &e, nil
		}
	}
	return nil, nil
}

func (m *mockService) UpdateAvatar(ctx context.Context, id int64, path string) error {
	return nil
}

func (m *mockService) DeleteEmployee(ctx context.Context, id int64) error {
	return nil
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	svc := &mockService{}
	ctrl := NewEmployeeController(svc)
	r.POST("/employees", ctrl.CreateEmployee)
	r.GET("/employees", ctrl.GetEmployees)
	return r
}

func TestCreateEmployeeHandler(t *testing.T) {
	r := setupTestRouter()

	emp := model.Employee{
		Name:        "Jane Doe",
		Email:       "jane@example.com",
		Department:  "HR",
		Designation: "Manager",
		Salary:      60000,
	}

	body, _ := json.Marshal(emp)
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	
	var res model.Employee
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, int64(1), res.ID)
	assert.Equal(t, "Jane Doe", res.Name)
}