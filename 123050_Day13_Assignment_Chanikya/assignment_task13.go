package main

import (
	"fmt"
	"net/http"
)

// Task 13: Liveness Probe Check Endpoint (/health)
// Simplistic endpoint asserting microservice runtime existence.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	fmt.Println("[Task 13] Liveness health probe setup logic defined.")
}
