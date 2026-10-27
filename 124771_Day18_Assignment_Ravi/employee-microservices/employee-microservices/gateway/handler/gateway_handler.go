package handler

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GatewayHandler forwards requests to the correct microservice using http.Client.
// This beginner-friendly approach is more explicit than httputil.ReverseProxy.
type GatewayHandler struct {
	client *http.Client
}

// NewGatewayHandler creates a gateway handler with a request timeout.
func NewGatewayHandler() *GatewayHandler {
	return &GatewayHandler{
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

// forwardRequest sends the incoming request to a target service and copies
// the response status, content type, and body back to the original client.
func (h *GatewayHandler) forwardRequest(c *gin.Context, targetURL string) {
	// Preserve query parameters such as ?page=1.
	if c.Request.URL.RawQuery != "" {
		targetURL += "?" + c.Request.URL.RawQuery
	}

	// Build a request using the incoming method and body (GET, POST, etc.).
	req, err := http.NewRequestWithContext(
		c.Request.Context(),
		c.Request.Method,
		targetURL,
		c.Request.Body,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create service request"})
		return
	}

	// Forward common headers needed by this beginner example.
	if contentType := c.GetHeader("Content-Type"); contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if accept := c.GetHeader("Accept"); accept != "" {
		req.Header.Set("Accept", accept)
	}

	// Send the request to the selected microservice.
	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "target microservice is unavailable"})
		return
	}
	defer resp.Body.Close()

	// Read the response returned by the microservice.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to read service response"})
		return
	}

	// Return the same status code and content type to the client.
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/json"
	}
	c.Data(resp.StatusCode, contentType, body)
}

// RegisterRoutes registers the public routes exposed by the API Gateway.
func (h *GatewayHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "api-gateway", "status": "ok"})
	})

	// Employee Service routes. Both GET and POST are forwarded.
	r.Any("/employees", func(c *gin.Context) {
		h.forwardRequest(c, "http://localhost:8081/employees")
	})

	// Department Service route.
	r.Any("/departments", func(c *gin.Context) {
		h.forwardRequest(c, "http://localhost:8082/departments")
	})

	// Auth Service route.
	r.Any("/login", func(c *gin.Context) {
		h.forwardRequest(c, "http://localhost:8083/login")
	})
}
