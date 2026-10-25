package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var ErrDepartmentNotFound = errors.New("department not found")
var ErrDepartmentUnavailable = errors.New("department service unavailable")

type DepartmentResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DepartmentClient struct {
	BaseURL string
	Client  *http.Client
	Breaker *CircuitBreaker
	Retries int
}

func NewDepartmentClient() *DepartmentClient {
	baseURL := strings.TrimRight(
		os.Getenv("DEPARTMENT_SERVICE_URL"),
		"/",
	)

	threshold, _ := strconv.Atoi(
		os.Getenv("CIRCUIT_BREAKER_FAILURE_THRESHOLD"),
	)

	resetSeconds, _ := strconv.Atoi(
		os.Getenv("CIRCUIT_BREAKER_RESET_TIMEOUT_SECONDS"),
	)

	return &DepartmentClient{
		BaseURL: baseURL,
		Client:  NewHTTPClient(),
		Breaker: NewCircuitBreaker(
			threshold,
			time.Duration(resetSeconds)*time.Second,
		),
		Retries: MaxRetries(),
	}
}

func (c *DepartmentClient) GetByID(
	ctx context.Context,
	id uint,
	correlationID string,
) (*DepartmentResponse, error) {
	if c.BaseURL == "" {
		return nil, fmt.Errorf(
			"%w: DEPARTMENT_SERVICE_URL is not configured",
			ErrDepartmentUnavailable,
		)
	}

	if id == 0 {
		return nil, fmt.Errorf("department ID must be greater than zero")
	}

	if err := c.Breaker.Allow(); err != nil {
		return nil, fmt.Errorf(
			"%w: %v",
			ErrDepartmentUnavailable,
			err,
		)
	}

	endpoint := c.BaseURL + "/api/departments/" +
		url.PathEscape(strconv.FormatUint(uint64(id), 10))

	resp, err := DoWithRetry(
		ctx,
		c.Client,
		func() (*http.Request, error) {
			req, err := http.NewRequest(
				http.MethodGet,
				endpoint,
				nil,
			)
			if err != nil {
				return nil, err
			}

			if correlationID != "" {
				req.Header.Set("X-Correlation-ID", correlationID)
			}

			return req, nil
		},
		c.Retries,
	)

	if err != nil {
		c.Breaker.RecordFailure()
		return nil, fmt.Errorf(
			"%w: %v",
			ErrDepartmentUnavailable,
			err,
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		c.Breaker.RecordSuccess()
		return nil, fmt.Errorf(
			"%w: ID %d",
			ErrDepartmentNotFound,
			id,
		)
	}

	if resp.StatusCode != http.StatusOK {
		c.Breaker.RecordFailure()
		return nil, fmt.Errorf(
			"%w: received HTTP %d",
			ErrDepartmentUnavailable,
			resp.StatusCode,
		)
	}

	var department DepartmentResponse
	if err := json.NewDecoder(resp.Body).Decode(&department); err != nil {
		c.Breaker.RecordFailure()
		return nil, fmt.Errorf(
			"%w: invalid response: %v",
			ErrDepartmentUnavailable,
			err,
		)
	}

	c.Breaker.RecordSuccess()
	return &department, nil
}
