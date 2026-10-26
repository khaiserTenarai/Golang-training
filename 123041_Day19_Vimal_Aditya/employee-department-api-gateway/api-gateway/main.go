package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"api-gateway/config"
	"api-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func createReverseProxy(targetBaseURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL := fmt.Sprintf("%s%s", targetBaseURL, c.Request.URL.Path)
		if c.Request.URL.RawQuery != "" {
			targetURL = fmt.Sprintf("%s?%s", targetURL, c.Request.URL.RawQuery)
		}

		req, err := http.NewRequestWithContext(c.Request.Context(), c.Request.Method, targetURL, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Copy request headers to target service
		for k, v := range c.Request.Header {
			req.Header[k] = v
		}

		// Ensure Correlation ID is attached
		corrID := c.GetString(middleware.CorrelationIDHeader)
		req.Header.Set(middleware.CorrelationIDHeader, corrID)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Target microservice unavailable: " + err.Error()})
			return
		}
		defer resp.Body.Close()

		// Forward response headers and body back to client
		for k, v := range resp.Header {
			c.Writer.Header()[k] = v
		}
		c.Writer.WriteHeader(resp.StatusCode)
		io.Copy(c.Writer, resp.Body)
	}
}

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()
	r.Use(middleware.CorrelationIDMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "api-gateway", "status": "UP"})
	})

	// Route employee requests to Employee Service (Port 8081)
	r.Any("/api/v1/employees", createReverseProxy(cfg.EmployeeServiceURL))
	r.Any("/api/v1/employees/*path", createReverseProxy(cfg.EmployeeServiceURL))

	// Route department requests to Department Service (Port 8082)
	r.Any("/api/v1/departments", createReverseProxy(cfg.DepartmentServiceURL))
	r.Any("/api/v1/departments/*path", createReverseProxy(cfg.DepartmentServiceURL))

	log.Printf("API Gateway running on port %s...", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
