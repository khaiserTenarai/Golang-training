package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// DepartmentClient calls Department Service over REST.
type DepartmentClient struct {
	baseURL string
	client  *http.Client
}

func NewDepartmentClient(baseURL string) *DepartmentClient {
	return &DepartmentClient{
		baseURL: strings.TrimRight(baseURL, "/"),
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

// Exists checks whether a department exists.
func (d *DepartmentClient) Exists(
	ctx context.Context,
	id uint,
) (bool, error) {
	url := fmt.Sprintf("%s/departments/%d", d.baseURL, id)

	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil,
	)
	if err != nil {
		return false, err
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return false, fmt.Errorf("department service unavailable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf(
			"department service returned status %d",
			resp.StatusCode,
		)
	}

	// Read the response to ensure it contains valid JSON.
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("invalid department response: %w", err)
	}

	return true, nil
}
