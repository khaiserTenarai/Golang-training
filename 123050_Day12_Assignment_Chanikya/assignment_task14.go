package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

// Task 14: Readiness state network probe endpoint (/ready)
// Interrogates external connections or local global lifecycle flags before routing live traffic.

var isReady int32 = 1 // 1 means application is ready

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&isReady) == 1 {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("READY"))
	} else {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
	}
}

func main() {
	fmt.Println("Task 14 - Readiness infrastructure logic instantiated.")
}
