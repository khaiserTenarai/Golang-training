package discovery

import (
	"fmt"
	"net/http"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	"order-management/config"
)

// NewConsulClient creates a client for the Consul agent API with a bounded timeout.
func NewConsulClient() (*consulapi.Client, error) {
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = config.Get("CONSUL_ADDRESS", "localhost:8500")
	consulConfig.Scheme = config.Get("CONSUL_SCHEME", "http")
	consulConfig.HttpClient = &http.Client{Timeout: config.GetDuration("CONSUL_TIMEOUT", time.Second)}
	client, err := consulapi.NewClient(consulConfig)
	if err != nil {
		return nil, fmt.Errorf("create Consul client: %w", err)
	}
	return client, nil
}
