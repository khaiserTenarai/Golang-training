package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"employee-department-microservices/employee-service/controller"
	"employee-department-microservices/employee-service/repository"
	"employee-department-microservices/employee-service/service"
	"employee-department-microservices/shared/utility"
	"github.com/gin-gonic/gin"
)

func TestEmployeeHealthContract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := repository.NewEmployeeRepository()
	svc := service.NewEmployeeService(repo, &http.Client{Timeout: time.Second}, "http://127.0.0.1:1", 1, utility.NewCircuitBreaker(2, time.Second))
	r := gin.New()
	r.Use(utility.CorrelationMiddleware())
	controller.NewEmployeeController(svc).Register(r)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/health", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if rec.Header().Get("X-Correlation-ID") == "" {
		t.Fatal("expected generated correlation ID")
	}
}
