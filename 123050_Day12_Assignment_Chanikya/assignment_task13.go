package main

import (
	"fmt"
	"net/http"
)

// Task 13: Liveness Probe check endpoint (/health)
// Simplistic endpoint determining whether the runtime runtime container process is up.

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func main() {
	fmt.Println("Task 13 - Liveness infrastructure logic instantiated.")
}
