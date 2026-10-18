package controller_test

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"employee-api/config"
	"employee-api/controller"
	"employee-api/middleware"
	"employee-api/repository"
	"employee-api/service"
)

const apiKey = "test-key-1234567890"

func newMux(exposeSalary bool) *http.ServeMux {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	opts := service.Options{RequestTimeout: time.Second, DefaultPageSize: 10, MaxPageSize: 100, SearchEnabled: true}
	svc := service.NewEmployeeService(repository.NewMemoryEmployeeRepository(), opts, logger)

	mux := http.NewServeMux()
	controller.NewEmployeeController(svc, logger, exposeSalary).Register(mux, middleware.APIKey(config.Secret(apiKey)))
	return mux
}

func do(mux http.Handler, method, path, body, key string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if key != "" {
		req.Header.Set("X-API-Key", key)
	}
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	return w
}

const newEmployee = `{"name":"Riya","email":"riya@example.com","department":"IT","salary":82000}`

func TestWritesNeedTheAPIKey(t *testing.T) {
	mux := newMux(true)
	if w := do(mux, "POST", "/employees", newEmployee, ""); w.Code != http.StatusUnauthorized {
		t.Fatalf("no key: want 401, got %d", w.Code)
	}
	if w := do(mux, "POST", "/employees", newEmployee, "wrong"); w.Code != http.StatusUnauthorized {
		t.Fatalf("wrong key: want 401, got %d", w.Code)
	}
	if w := do(mux, "POST", "/employees", newEmployee, apiKey); w.Code != http.StatusCreated {
		t.Fatalf("right key: want 201, got %d %s", w.Code, w.Body.String())
	}
	if w := do(mux, "GET", "/employees", "", ""); w.Code != http.StatusOK {
		t.Fatalf("reading is public: want 200, got %d", w.Code)
	}
}

func TestSalaryFeatureFlag(t *testing.T) {
	shown, hidden := newMux(true), newMux(false)
	do(shown, "POST", "/employees", newEmployee, apiKey)
	do(hidden, "POST", "/employees", newEmployee, apiKey)

	if body := do(shown, "GET", "/employees/1", "", "").Body.String(); !strings.Contains(body, `"salary":82000`) {
		t.Fatalf("salary should be visible: %s", body)
	}
	if body := do(hidden, "GET", "/employees/1", "", "").Body.String(); strings.Contains(body, "salary") {
		t.Fatalf("salary must be hidden: %s", body)
	}
}

func TestErrorsMapToStatusCodes(t *testing.T) {
	mux := newMux(true)
	if w := do(mux, "POST", "/employees", `{"name":"R","email":"x"}`, apiKey); w.Code != http.StatusBadRequest {
		t.Fatalf("invalid body: want 400, got %d", w.Code)
	}
	if w := do(mux, "GET", "/employees/99", "", ""); w.Code != http.StatusNotFound {
		t.Fatalf("missing employee: want 404, got %d", w.Code)
	}
	if w := do(mux, "GET", "/employees/abc", "", ""); w.Code != http.StatusBadRequest {
		t.Fatalf("bad id: want 400, got %d", w.Code)
	}
	do(mux, "POST", "/employees", newEmployee, apiKey)
	if w := do(mux, "POST", "/employees", newEmployee, apiKey); w.Code != http.StatusConflict {
		t.Fatalf("duplicate email: want 409, got %d", w.Code)
	}
	if w := do(mux, "DELETE", "/employees/1", "", apiKey); w.Code != http.StatusNoContent {
		t.Fatalf("delete: want 204, got %d", w.Code)
	}
}
