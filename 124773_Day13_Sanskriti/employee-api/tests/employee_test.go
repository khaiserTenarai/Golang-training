package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"employee-api/internal/controller"
	"employee-api/internal/model"
	"employee-api/internal/repository"
	"employee-api/internal/server"
	"employee-api/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeAPI_CreateAndGet(t *testing.T) {
	repo := repository.NewEmployeeRepository(nil)
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc)
	router := server.NewRouter(ctrl)

	emp := model.Employee{
		ID:       "emp-101",
		Name:     "Jane Doe",
		Email:    "jane.doe@example.com",
		Position: "Software Engineer",
	}
	body, _ := json.Marshal(emp)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "POST", "/api/v1/employees", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	wFetch := httptest.NewRecorder()
	reqFetch, _ := http.NewRequestWithContext(context.Background(), "GET", "/api/v1/employees/emp-101", nil)
	router.ServeHTTP(wFetch, reqFetch)

	assert.Equal(t, http.StatusOK, wFetch.Code)

	var fetchedEmp model.Employee
	_ = json.Unmarshal(wFetch.Body.Bytes(), &fetchedEmp)
	assert.Equal(t, emp.Name, fetchedEmp.Name)
}
