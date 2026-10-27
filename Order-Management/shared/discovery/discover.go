package discovery

import (
	"fmt"
	"net"
	"net/url"

	consulapi "github.com/hashicorp/consul/api"
)

// Discover finds healthy instances of a registered service.
func Discover(
	client *consulapi.Client,
	serviceName string,
) ([]string, error) {
	if client == nil {
		return nil, fmt.Errorf("Consul client is nil")
	}

	instances, _, err := client.Health().Service(
		serviceName,
		"",
		true,
		&consulapi.QueryOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"discover service %s: %w",
			serviceName,
			err,
		)
	}

	var addresses []string

	for _, instance := range instances {
		address := instance.Service.Address
		if address == "" {
			address = instance.Node.Address
		}

		if address == "" || instance.Service.Port == 0 {
			continue
		}

		serviceURL := url.URL{
			Scheme: "http",
			Host: net.JoinHostPort(
				address,
				fmt.Sprintf("%d", instance.Service.Port),
			),
		}

		addresses = append(addresses, serviceURL.String())
	}

	if len(addresses) == 0 {
		return nil, fmt.Errorf(
			"no healthy instances found for service %s",
			serviceName,
		)
	}

	return addresses, nil
}

// DiscoverOne returns the address of one healthy service instance.
func DiscoverOne(
	client *consulapi.Client,
	serviceName string,
) (string, error) {
	addresses, err := Discover(client, serviceName)
	if err != nil {
		return "", err
	}

	return addresses[0], nil
}
