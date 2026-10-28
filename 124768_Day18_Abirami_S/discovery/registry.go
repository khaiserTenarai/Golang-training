package discovery

import "sync"

// Registry is a small in-memory demo registry. It is not a health-aware,
// distributed registry like Consul or Kubernetes Service Discovery.
type Registry struct {
    mu sync.RWMutex
    services map[string]string
}

func NewRegistry() *Registry {
    return &Registry{services: make(map[string]string)}
}

func (r *Registry) Register(name, address string) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.services[name] = address
}

func (r *Registry) Resolve(name string) (string, bool) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    address, ok := r.services[name]
    return address, ok
}
