package main

import (
	"fmt"
	"time"
)

// Task 6: Operational Microsecond Logging Middleware
// Measures exact processing time of arbitrary functions.
func TelemetryWrapper(action func()) {
	start := time.Now()
	action()
	duration := time.Since(start).Microseconds()
	fmt.Printf("[Task 6] Action execution telemetry: %d microseconds\n", duration)
}

func main() {
	TelemetryWrapper(func() {
		// Simulate arbitrary execution work
		time.Sleep(2 * time.Millisecond)
	})
}
