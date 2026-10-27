package discovery

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

// RoundRobin resolves healthy Consul instances and selects them in round-robin order.
// It caches discovery results briefly to avoid querying Consul on every HTTP request.
type RoundRobin struct {
	client     *consulapi.Client
	name       string
	fallback   string
	cacheTTL   time.Duration
	mu         sync.Mutex
	cached     []string
	cacheUntil time.Time
	next       atomic.Uint64
}

func NewRoundRobin(client *consulapi.Client, serviceName, fallback string) *RoundRobin {
	return &RoundRobin{client: client, name: serviceName, fallback: fallback, cacheTTL: 3 * time.Second}
}

func (b *RoundRobin) Next() (*url.URL, error) {
	instances, err := b.instances()
	if err != nil {
		if b.fallback != "" {
			return url.Parse(b.fallback)
		}
		return nil, err
	}
	idx := b.next.Add(1) - 1
	return url.Parse(instances[idx%uint64(len(instances))])
}

func (b *RoundRobin) instances() ([]string, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if time.Now().Before(b.cacheUntil) && len(b.cached) > 0 {
		return append([]string(nil), b.cached...), nil
	}
	if b.client == nil {
		if b.fallback != "" {
			return []string{b.fallback}, nil
		}
		return nil, fmt.Errorf("Consul client is unavailable")
	}
	items, _, err := b.client.Health().Service(b.name, "", true, &consulapi.QueryOptions{})
	if err != nil {
		if len(b.cached) > 0 {
			return append([]string(nil), b.cached...), nil
		}
		if b.fallback != "" {
			return []string{b.fallback}, nil
		}
		return nil, err
	}
	var addresses []string
	for _, item := range items {
		host := item.Service.Address
		if host == "" {
			host = item.Node.Address
		}
		if host != "" && item.Service.Port > 0 {
			addresses = append(addresses, fmt.Sprintf("http://%s:%d", host, item.Service.Port))
		}
	}
	if len(addresses) == 0 {
		if b.fallback != "" {
			return []string{b.fallback}, nil
		}
		return nil, fmt.Errorf("no healthy instances for %s", b.name)
	}
	b.cached = append([]string(nil), addresses...)
	b.cacheUntil = time.Now().Add(b.cacheTTL)
	return append([]string(nil), addresses...), nil
}
