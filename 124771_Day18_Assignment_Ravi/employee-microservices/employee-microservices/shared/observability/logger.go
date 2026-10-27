package observability

import (
	"log"
	"os"
)

// NewLogger provides a named standard-library logger. Task 14: basic logs.
// Metrics and distributed tracing are intentionally left for a later learning step.
func NewLogger(serviceName string) *log.Logger {
	return log.New(os.Stdout, "["+serviceName+"] ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
}
