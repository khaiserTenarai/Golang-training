package discovery

import (
	"fmt"
	"strings"

	"order-management/config"
	consulapi "github.com/hashicorp/consul/api"
)

// Register registers a service and its HTTP health check with Consul.
func Register(
	client *consulapi.Client,
	serviceName string,
	serviceID string,
	address string,
	port int,
	healthPath string,
	interval string,
	timeout string,
) error {
	if client == nil {
		return fmt.Errorf("Consul client is nil")
	}

	if serviceName == "" || serviceID == "" || address == "" {
		return fmt.Errorf("service name, ID, and address are required")
	}

	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid service port: %d", port)
	}

	if !strings.HasPrefix(healthPath, "/") {
		healthPath = "/" + healthPath
	}

	healthURL := fmt.Sprintf(
		"http://%s:%d%s",
		address,
		port,
		healthPath,
	)

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: address,
		Port:    port,
		Check: &consulapi.AgentServiceCheck{
			HTTP:                           healthURL,
			Interval:                       interval,
			Timeout:                        timeout,
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		return fmt.Errorf(
			"register %s with Consul: %w",
			serviceName,
			err,
		)
	}

	return nil
}

// Deregister removes a service registration from Consul.
func Deregister(
	client *consulapi.Client,
	serviceID string,
) error {
	if client == nil {
		return fmt.Errorf("Consul client is nil")
	}

	if err := client.Agent().ServiceDeregister(serviceID); err != nil {
		return fmt.Errorf(
			"deregister service %s: %w",
			serviceID,
			err,
		)
	}

	return nil
}

// RegisterFromEnv registers a local service using the common environment settings.
func RegisterFromEnv(client *consulapi.Client, serviceName string, port int, healthPath string) (string, error) {
	address := config.Get("SERVICE_ADVERTISE_ADDRESS", "127.0.0.1")
	id := fmt.Sprintf("%s-%d", serviceName, port)
	return id, Register(client, serviceName, id, address, port, healthPath,
		config.Get("HEALTH_CHECK_INTERVAL", "10s"),
		config.Get("HEALTH_CHECK_TIMEOUT", "3s"))
}
