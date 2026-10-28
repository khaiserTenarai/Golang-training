package main

import (
	"employee-department-microservices/employee-service/config"
	"employee-department-microservices/employee-service/controller"
	"employee-department-microservices/employee-service/repository"
	"employee-department-microservices/employee-service/service"
	"employee-department-microservices/shared/utility"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnv(".env")
	cfg := config.Load()
	repo := repository.NewEmployeeRepository()
	breaker := utility.NewCircuitBreaker(cfg.BreakerThreshold, cfg.BreakerOpenTimeout)
	client := &http.Client{Timeout: cfg.HTTPTimeout}
	svc := service.NewEmployeeService(repo, client, cfg.DepartmentURL, cfg.RetryCount, breaker)
	r := gin.New()
	r.Use(gin.Recovery(), utility.CorrelationMiddleware(), utility.LoggingMiddleware())
	controller.NewEmployeeController(svc).Register(r)
	log.Printf("employee service listening on :%s; department_url=%s", cfg.Port, cfg.DepartmentURL)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
func loadEnv(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	for _, line := range splitLines(string(b)) {
		line = trim(line)
		if line == "" || line[0] == '#' {
			continue
		}
		k, v, ok := cut(line, "=")
		if ok && os.Getenv(trim(k)) == "" {
			_ = os.Setenv(trim(k), trim(v))
		}
	}
}
func splitLines(s string) []string {
	var out []string
	start := 0
	for i, c := range s {
		if c == '\n' {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		out = append(out, s[start:])
	}
	return out
}
func trim(s string) string {
	for len(s) > 0 && (s[0] == ' ' || s[0] == '\t' || s[0] == '\r') {
		s = s[1:]
	}
	for len(s) > 0 && (s[len(s)-1] == ' ' || s[len(s)-1] == '\t' || s[len(s)-1] == '\r') {
		s = s[:len(s)-1]
	}
	return s
}
func cut(s, sep string) (string, string, bool) {
	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			return s[:i], s[i+len(sep):], true
		}
	}
	return "", "", false
}
