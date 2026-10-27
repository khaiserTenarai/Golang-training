package httpclient

import (
	"net"
	"net/http"
	"time"

	"order-management/config"
)

// New creates an HTTP client with a configurable timeout.
func New() *http.Client {
	timeout := config.GetDuration(
		"HTTP_CLIENT_TIMEOUT",
		10*time.Second,
	)

	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
	}
}
