package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"employee-management-microservices/shared/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func serviceProxy(target string) gin.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil || targetURL.Scheme == "" || targetURL.Host == "" {
		log.Fatalf("Invalid upstream service URL: %s", target)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ErrorHandler = func(
		w http.ResponseWriter,
		r *http.Request,
		err error,
	) {
		log.Printf("Upstream service error: %v", err)
		http.Error(
			w,
			`{"error":"Upstream service unavailable"}`,
			http.StatusBadGateway,
		)
	}

	return func(c *gin.Context) {
		// Preserve the original path, including /api.
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	_ = godotenv.Load(".env")

	employeeURL := os.Getenv("EMPLOYEE_SERVICE_URL")
	if employeeURL == "" {
		employeeURL = "http://localhost:8081"
	}

	departmentURL := os.Getenv("DEPARTMENT_SERVICE_URL")
	if departmentURL == "" {
		departmentURL = "http://localhost:8082"
	}

	port := os.Getenv("GATEWAY_PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.Use(middleware.CorrelationID())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "UP",
			"service": "gateway",
		})
	})

	employeeProxy := serviceProxy(employeeURL)
	departmentProxy := serviceProxy(departmentURL)

	router.Any("/api/employees", employeeProxy)
	router.Any("/api/employees/*path", employeeProxy)

	router.Any("/api/departments", departmentProxy)
	router.Any("/api/departments/*path", departmentProxy)

	log.Printf("API Gateway running on port %s", port)
	log.Fatal(router.Run(":" + port))
}
