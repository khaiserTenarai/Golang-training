package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)


func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}


func sendRequest(router http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}

func TestCreateEmployee(t *testing.T) {
	router := NewRouter()

	rec := sendRequest(router, "POST", "/employees", `{"name":"Anita","salary":45000}`)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}

	var got Employee
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("response was not valid JSON: %v", err)
	}
	if got.ID != 1 || got.Name != "Anita" {
		t.Errorf("unexpected employee in response: %+v", got)
	}
}

func TestCreateEmployeeInvalid(t *testing.T) {
	router := NewRouter()

	rec := sendRequest(router, "POST", "/employees", `{"name":"","salary":-5}`)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}

func TestGetEmployee(t *testing.T) {
	router := NewRouter()
	sendRequest(router, "POST", "/employees", `{"name":"Ravi","salary":60000}`)

	rec := sendRequest(router, "GET", "/employees/1", "")

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Ravi") {
		t.Errorf("expected the response to contain Ravi, got %s", rec.Body.String())
	}
}

func TestGetEmployeeNotFound(t *testing.T) {
	router := NewRouter()

	rec := sendRequest(router, "GET", "/employees/99", "")

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rec.Code)
	}
}

func TestGetEmployeeBadID(t *testing.T) {
	router := NewRouter()

	rec := sendRequest(router, "GET", "/employees/abc", "")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}
