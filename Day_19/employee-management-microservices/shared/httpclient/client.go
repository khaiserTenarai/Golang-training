package httpclient

import (
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

// NewHTTPClient creates a shared HTTP client for communication
// between microservices.
func NewHTTPClient() *http.Client {
	// Default maximum duration for an HTTP request is 5 seconds.
	timeoutSeconds := 5

	// Allow the timeout to be configured through an environment variable.
	if value := os.Getenv("HTTP_CLIENT_TIMEOUT_SECONDS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			timeoutSeconds = parsed
		}
	}

	return &http.Client{
		// Stop waiting when the total request duration reaches this limit.
		Timeout: time.Duration(timeoutSeconds) * time.Second,

		Transport: &http.Transport{
			// Maximum time allowed to establish a network connection.
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,

			// Connection pooling improves performance when connections are reused.
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,

			// Close idle connections after 90 seconds.
			IdleConnTimeout: 90 * time.Second,
		},
	}
}
