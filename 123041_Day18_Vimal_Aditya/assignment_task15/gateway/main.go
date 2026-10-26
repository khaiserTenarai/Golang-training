package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"day18/config"

	"github.com/gin-gonic/gin"
)

func createGinProxy(targetURL string) gin.HandlerFunc {
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("Invalid target URL %s: %v", targetURL, err)
	}
	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	return func(c *gin.Context) {
		c.Request.Host = parsedURL.Host
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()

	employeeProxy := createGinProxy("http://localhost:" + cfg.AppPort)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "api-gateway", "status": "UP"})
	})

	r.Any("/metrics", employeeProxy)
	r.Any("/graphql", employeeProxy)
	r.Any("/api/v1/employees", employeeProxy)
	r.Any("/api/v1/employees/*path", employeeProxy)

	gatewayPort := "8000"
	fmt.Printf("API Gateway listening on http://localhost:%s\n", gatewayPort)
	if err := r.Run(":" + gatewayPort); err != nil {
		log.Fatalf("API Gateway crashed: %v", err)
	}
}