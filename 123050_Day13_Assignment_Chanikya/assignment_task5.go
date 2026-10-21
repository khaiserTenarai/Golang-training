package main

import (
	"context"
	"fmt"
)

// Task 5: Request-ID Generating Middleware Using Context
// Decorates a request lifecycle context with unique tracking credentials.
type ctxKey string

const RequestIDKey ctxKey = "request_id"

func main() {
	ctx := context.Background()
	// Simulate middleware assignment
	reqCtx := context.WithValue(ctx, RequestIDKey, "REQ-TX-998822")

	// Simulate retrieving down the stack
	if rid, ok := reqCtx.Value(RequestIDKey).(string); ok {
		fmt.Printf("[Task 5] Extracted Request ID from Context: %s\n", rid)
	}
}
