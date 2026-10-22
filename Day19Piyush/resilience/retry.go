// Task 11: Retry with exponential backoff.

package resilience

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type RetryConfig struct {
	MaxAttempts int
	BaseDelay   time.Duration
}

func DefaultRetryConfig() RetryConfig {
	return RetryConfig{MaxAttempts: 3, BaseDelay: 200 * time.Millisecond}
}

func DoWithRetry(client *http.Client, req *http.Request, cfg RetryConfig) (*http.Response, error) {
	var lastErr error
	for attempt := 1; attempt <= cfg.MaxAttempts; attempt++ {
		resp, err := client.Do(req)
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}
		if err != nil {
			lastErr = err
		} else {
			lastErr = fmt.Errorf("server returned %d", resp.StatusCode)
			resp.Body.Close()
		}
		if attempt < cfg.MaxAttempts {
			delay := cfg.BaseDelay * time.Duration(1<<(attempt-1))
			log.Printf("[Retry] Attempt %d/%d failed: %v. Retrying in %v...",
				attempt, cfg.MaxAttempts, lastErr, delay)
			time.Sleep(delay)
		}
	}
	return nil, fmt.Errorf("all %d attempts failed: %w", cfg.MaxAttempts, lastErr)
}
