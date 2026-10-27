package discovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// RegisterWithConsul demonstrates optional Consul service registration. Task 11.
// It requires a local Consul agent, e.g. at http://localhost:8500.
// This function is not called by default, so the basic project runs without Consul.
func RegisterWithConsul(consulAddress, serviceID, serviceName, address string, port int) error {
	body := map[string]any{
		"ID": serviceID, "Name": serviceName, "Address": address, "Port": port,
		"Check": map[string]any{"HTTP": fmt.Sprintf("http://%s:%d/health", address, port), "Interval": "10s", "Timeout": "2s"},
	}
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, consulAddress+"/v1/agent/service/register", bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Consul registration returned %s", resp.Status)
	}
	return nil
}
