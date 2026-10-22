// Tasks 4, 9, 10: HTTP Client for service-to-service communication.

package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go_microservices/dto"
	"go_microservices/middleware"
	"go_microservices/resilience"
)

type ServiceClient struct {
	client  *http.Client
	baseURL string
	cb      *resilience.CircuitBreaker
	retry   resilience.RetryConfig
}

func NewServiceClient(baseURL string, timeoutSecs int) *ServiceClient {
	return &ServiceClient{
		client:  &http.Client{Timeout: time.Duration(timeoutSecs) * time.Second},
		baseURL: baseURL,
		cb:      resilience.NewCircuitBreaker(baseURL, 3),
		retry:   resilience.DefaultRetryConfig(),
	}
}

func (sc *ServiceClient) Get(ctx context.Context, path string) ([]byte, error) {
	url := sc.baseURL + path
	var body []byte
	err := sc.cb.Execute(func() error {
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return fmt.Errorf("create request: %w", err)
		}
		if cid := middleware.GetCorrelationID(ctx); cid != "" {
			req.Header.Set("X-Correlation-ID", cid)
		}
		resp, err := resilience.DoWithRetry(sc.client, req, sc.retry)
		if err != nil {
			return fmt.Errorf("service call failed: %w", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			respBody, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("downstream error %d: %s", resp.StatusCode, string(respBody))
		}
		body, err = io.ReadAll(resp.Body)
		return err
	})
	return body, err
}

func (sc *ServiceClient) GetDepartment(ctx context.Context, id int) *dto.DepartmentResponse {
	data, err := sc.Get(ctx, fmt.Sprintf("/departments/%d", id))
	if err != nil {
		log.Printf("[Client] Failed to fetch department %d: %v", id, err)
		return nil
	}
	var resp dto.APIResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil
	}
	if resp.Data == nil {
		return nil
	}
	deptBytes, _ := json.Marshal(resp.Data)
	var dept dto.DepartmentResponse
	json.Unmarshal(deptBytes, &dept)
	return &dept
}

func (sc *ServiceClient) CircuitState() string {
	return sc.cb.GetState()
}
