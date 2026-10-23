package contract

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	deptController "employee-department-microservices/department-service/controller"
	deptRepo "employee-department-microservices/department-service/repository"
	deptService "employee-department-microservices/department-service/service"
	"employee-department-microservices/shared/utility"
	"github.com/gin-gonic/gin"
)

func TestDepartmentHealthContract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(utility.CorrelationMiddleware())
	deptController.NewDepartmentController(deptService.NewDepartmentService(deptRepo.NewDepartmentRepository())).Register(r)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set("X-Correlation-ID", "contract-test")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body["service"] != "department" || body["status"] != "UP" {
		t.Fatalf("unexpected health response: %v", body)
	}
	if rec.Header().Get("X-Correlation-ID") != "contract-test" {
		t.Fatal("correlation ID was not echoed")
	}
}

func TestDepartmentListContract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(utility.CorrelationMiddleware())
	deptController.NewDepartmentController(deptService.NewDepartmentService(deptRepo.NewDepartmentRepository())).Register(r)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/departments", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var body []map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if len(body) < 1 {
		t.Fatal("expected seeded departments")
	}
	for _, field := range []string{"id", "name", "description"} {
		if _, ok := body[0][field]; !ok {
			t.Fatalf("missing contract field %q", field)
		}
	}
}
