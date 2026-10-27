package client

import (
	"context"
	"day19/employee-service/dto"
	"day19/employee-service/resilience"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const correlationHeader = "X-Correlation-ID"

// DepartmentClient is the only layer that knows the remote service URL/protocol.
type DepartmentClient struct {
	baseURL    string
	httpClient *http.Client // Task 10: timeout prevents hanging forever.
	retries    int
	breaker    *resilience.CircuitBreaker
}

func NewDepartmentClient(baseURL string, timeout time.Duration, retries int, breaker *resilience.CircuitBreaker) *DepartmentClient {
	return &DepartmentClient{baseURL: strings.TrimRight(baseURL, "/"), httpClient: &http.Client{Timeout: timeout}, retries: retries, breaker: breaker}
}
func (c *DepartmentClient) Get(ctx context.Context, id uint, correlationID string) (*dto.DepartmentDTO, error) {
	// Task 13: do not keep calling a failing downstream service while breaker is open.
	if err := c.breaker.Allow(); err != nil {
		return nil, err
	}
	var lastErr error
	// Task 11: retry a bounded number of times for transport errors and 5xx only.
	for attempt := 0; attempt <= c.retries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/departments/%d", c.baseURL, id), nil)
		if err != nil {
			return nil, err
		}
		// Task 12: propagate correlation ID to the downstream service.
		if correlationID != "" {
			req.Header.Set(correlationHeader, correlationID)
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt < c.retries {
				time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
				continue
			}
			c.breaker.Failure()
			return nil, fmt.Errorf("department service request failed: %w", lastErr)
		}
		body, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			lastErr = readErr
			c.breaker.Failure()
			return nil, readErr
		}
		if resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("department service returned HTTP %d", resp.StatusCode)
			if attempt < c.retries {
				time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
				continue
			}
			c.breaker.Failure()
			return nil, lastErr
		}
		if resp.StatusCode == http.StatusNotFound {
			c.breaker.Success()
			return nil, fmt.Errorf("department %d not found", id)
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			c.breaker.Success()
			return nil, fmt.Errorf("department service returned HTTP %d: %s", resp.StatusCode, string(body))
		}
		var out dto.DepartmentDTO
		if err := json.Unmarshal(body, &out); err != nil {
			c.breaker.Failure()
			return nil, fmt.Errorf("decode department response: %w", err)
		}
		c.breaker.Success()
		return &out, nil
	}
	c.breaker.Failure()
	return nil, lastErr
}
