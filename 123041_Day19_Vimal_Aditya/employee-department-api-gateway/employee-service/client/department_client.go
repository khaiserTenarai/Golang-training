package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"employee-service/dto"

	"github.com/sony/gobreaker"
)

type DepartmentClient interface {
	GetDepartmentByID(ctx context.Context, id int64, correlationID string) (*dto.DepartmentDTO, error)
}

type departmentClient struct {
	baseURL string
	client  *http.Client
	cb      *gobreaker.CircuitBreaker
}

func NewDepartmentClient(baseURL string) DepartmentClient {
	st := gobreaker.Settings{
		Name:        "department-service",
		MaxRequests: 1,
		Interval:    10 * time.Second,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
	}
	return &departmentClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 1000 * time.Millisecond},
		cb:      gobreaker.NewCircuitBreaker(st),
	}
}

func (c *departmentClient) GetDepartmentByID(ctx context.Context, id int64, correlationID string) (*dto.DepartmentDTO, error) {
	result, err := c.cb.Execute(func() (interface{}, error) {
		var lastErr error
		url := fmt.Sprintf("%s/api/v1/departments/%d", c.baseURL, id)

		for attempt := 1; attempt <= 3; attempt++ {
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return nil, err
			}
			req.Header.Set("X-Correlation-ID", correlationID)

			resp, err := c.client.Do(req)
			if err == nil {
				defer resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
					var dept dto.DepartmentDTO
					if err := json.NewDecoder(resp.Body).Decode(&dept); err != nil {
						return nil, err
					}
					return &dept, nil
				}
				lastErr = fmt.Errorf("department service returned HTTP %d", resp.StatusCode)
			} else {
				lastErr = err
			}
			time.Sleep(time.Duration(100*(1<<(attempt-1))) * time.Millisecond)
		}
		return nil, lastErr
	})

	if err != nil {
		return nil, err
	}
	return result.(*dto.DepartmentDTO), nil
}
