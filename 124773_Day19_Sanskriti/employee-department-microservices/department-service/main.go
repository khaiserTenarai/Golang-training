package main

import (
	"employee-department-microservices/department-service/config"
	"employee-department-microservices/department-service/controller"
	"employee-department-microservices/department-service/repository"
	"employee-department-microservices/department-service/service"
	"employee-department-microservices/shared/utility"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	loadEnv(".env")
	r := gin.New()
	r.Use(gin.Recovery(), utility.CorrelationMiddleware(), utility.LoggingMiddleware())
	repo := repository.NewDepartmentRepository()
	svc := service.NewDepartmentService(repo)
	controller.NewDepartmentController(svc).Register(r)
	port := config.Port()
	log.Printf("department service listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
func loadEnv(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	for _, line := range lines(string(b)) {
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
func lines(s string) []string {
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
