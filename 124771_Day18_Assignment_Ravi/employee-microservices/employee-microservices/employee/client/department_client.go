package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// DepartmentClient demonstrates synchronous REST from Employee to Department Service.
// Task 8. The client calls a different local service over HTTP.
type DepartmentClient struct {
	BaseURL string
	HTTP    *http.Client
}

func NewDepartmentClient(baseURL string) *DepartmentClient {
	return &DepartmentClient{BaseURL: baseURL, HTTP: &http.Client{Timeout: 3 * time.Second}}
}

func (c *DepartmentClient) ListDepartments() (map[string]any, error) {
	resp, err := c.HTTP.Get(c.BaseURL + "/departments")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("department service returned %s", resp.Status)
	}
	var data map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
