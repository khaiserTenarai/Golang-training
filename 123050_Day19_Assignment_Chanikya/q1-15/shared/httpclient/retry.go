package httpclient

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func MaxRetries() int {
	retries := 3

	if value := os.Getenv("HTTP_MAX_RETRIES"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed >= 0 {
			retries = parsed
		}
	}

	return retries
}

// DoWithRetry retries transient network errors and HTTP 5xx responses.
// The request must be safe to retry and have a replayable body.
func DoWithRetry(
	ctx context.Context,
	client *http.Client,
	newRequest func() (*http.Request, error),
	maxRetries int,
) (*http.Response, error) {
	if maxRetries < 0 {
		maxRetries = 0
	}

	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		req, err := newRequest()
		if err != nil {
			return nil, err
		}

		req = req.WithContext(ctx)

		resp, err := client.Do(req)
		if err == nil && (resp.StatusCode < 500 || resp.StatusCode > 599) {
			return resp, nil
		}

		if err == nil {
			lastErr = fmt.Errorf("temporary server error: HTTP %d", resp.StatusCode)
			resp.Body.Close()
		} else {
			lastErr = err
			if !isRetryableNetworkError(err) {
				return nil, err
			}
		}

		if attempt == maxRetries {
			break
		}

		delay := time.Duration(attempt+1) * 200 * time.Millisecond
		timer := time.NewTimer(delay)

		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}

	return nil, fmt.Errorf("request failed after retries: %w", lastErr)
}

func isRetryableNetworkError(err error) bool {
	if err == nil {
		return false
	}

	if netErr, ok := err.(net.Error); ok {
		return netErr.Timeout() || netErr.Temporary()
	}

	return true
}
