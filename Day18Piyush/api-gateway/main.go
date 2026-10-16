// Task 10: Design API Gateway
// Task 11: Service Discovery (registry-based)
// Task 8: Synchronous REST — gateway proxies requests to services
// Run: go run main.go
// Port: 8080 (single entry point)
//
// Routes all requests through one port to the correct microservice.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Task 11: Service Registry (simple in-memory discovery)
type ServiceInfo struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Healthy bool   `json:"healthy"`
}

var serviceRegistry = map[string]ServiceInfo{
	"employee":   {Name: "employee-service", URL: "http://localhost:8081"},
	"department": {Name: "department-service", URL: "http://localhost:8082"},
	"auth":       {Name: "auth-service", URL: "http://localhost:8083"},
}

func respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// Proxy a request to a backend service
func proxy(w http.ResponseWriter, r *http.Request, targetURL string) {
	client := &http.Client{Timeout: 5 * time.Second}
	proxyURL := targetURL

	req, err := http.NewRequest(r.Method, proxyURL, r.Body)
	if err != nil {
		respond(w, 502, map[string]interface{}{"error": "failed to create request"})
		return
	}
	req.Header = r.Header.Clone()

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[Gateway] Service unavailable: %s → %v", targetURL, err)
		respond(w, 503, map[string]interface{}{"error": "service unavailable", "target": targetURL})
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Proxied-By", "api-gateway")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// Auth validation — calls auth-service to verify token
func validateToken(r *http.Request) (bool, string) {
	if r.Header.Get("Authorization") == "" {
		return false, ""
	}
	client := &http.Client{Timeout: 2 * time.Second}
	req, _ := http.NewRequest("GET", serviceRegistry["auth"].URL+"/auth/validate", nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return false, ""
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	if valid, ok := result["valid"].(bool); ok && valid {
		role, _ := result["role"].(string)
		return true, role
	}
	return false, ""
}

func gatewayHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Printf("[Gateway] %s %s", r.Method, path)

	// Public routes — no auth needed
	if strings.HasPrefix(path, "/auth/") {
		proxy(w, r, serviceRegistry["auth"].URL+path)
		return
	}
	if path == "/health" || path == "/services" {
		if path == "/services" {
			respond(w, 200, serviceRegistry)
			return
		}
		respond(w, 200, map[string]interface{}{"service": "api-gateway", "status": "healthy"})
		return
	}

	// Protected routes — validate token via auth-service
	valid, role := validateToken(r)
	if !valid {
		respond(w, 401, map[string]interface{}{"error": "unauthorized — login via POST /auth/login"})
		return
	}

	// Write operations need admin/manager
	if r.Method != "GET" && role == "user" {
		respond(w, 403, map[string]interface{}{"error": "write access requires admin or manager role"})
		return
	}

	// Route to correct service
	switch {
	case strings.HasPrefix(path, "/employees"):
		proxy(w, r, serviceRegistry["employee"].URL+path)
	case strings.HasPrefix(path, "/departments"):
		proxy(w, r, serviceRegistry["department"].URL+path)
	default:
		respond(w, 404, map[string]interface{}{"error": "route not found"})
	}
}

// Task 11: Health check all services
func checkHealth() {
	for name, svc := range serviceRegistry {
		client := &http.Client{Timeout: 2 * time.Second}
		resp, err := client.Get(svc.URL + "/health")
		svc.Healthy = err == nil && resp != nil && resp.StatusCode == 200
		serviceRegistry[name] = svc
		status := "DOWN"
		if svc.Healthy {
			status = "UP"
		}
		fmt.Printf("  %s (%s): %s\n", svc.Name, svc.URL, status)
	}
}

func main() {
	fmt.Println("[API Gateway] Checking service health...")
	checkHealth()

	fmt.Println("\n[API Gateway] Running on http://localhost:8080")
	fmt.Println("Routes:")
	fmt.Println("  POST /auth/login           → auth-service (public)")
	fmt.Println("  GET  /auth/validate         → auth-service")
	fmt.Println("  GET  /employees             → employee-service (auth required)")
	fmt.Println("  POST /employees             → employee-service (admin/manager)")
	fmt.Println("  GET  /departments           → department-service (auth required)")
	fmt.Println("  GET  /services              → service registry")
	fmt.Println("  GET  /health                → gateway health")

	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(gatewayHandler)))
}
