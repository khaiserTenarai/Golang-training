package resilience

import (
	"errors"
	"sync"
	"time"
)

var ErrCircuitOpen = errors.New("department service circuit breaker is open")

// CircuitBreaker is a small in-memory educational breaker. After enough
// consecutive failures it rejects calls for openFor; the next call probes recovery.
type CircuitBreaker struct {
	mu        sync.Mutex
	failures  int
	threshold int
	openFor   time.Duration
	openedAt  time.Time
}

func NewCircuitBreaker(threshold int, openFor time.Duration) *CircuitBreaker {
	return &CircuitBreaker{threshold: threshold, openFor: openFor}
}
func (b *CircuitBreaker) Allow() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.failures < b.threshold {
		return nil
	}
	if time.Since(b.openedAt) < b.openFor {
		return ErrCircuitOpen
	}
	// Permit a probe after the cooldown; subsequent success/failure updates the state.
	b.failures = b.threshold - 1
	b.openedAt = time.Time{} // clear old open timestamp; a failed probe can open it again
	return nil
}
func (b *CircuitBreaker) Success() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.failures = 0
	b.openedAt = time.Time{}
}
func (b *CircuitBreaker) Failure() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.failures++
	if b.failures >= b.threshold && b.openedAt.IsZero() {
		b.openedAt = time.Now()
	}
}
