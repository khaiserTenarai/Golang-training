package utility

import (
	"errors"
	"sync"
	"time"
)

var ErrCircuitOpen = errors.New("department service circuit breaker is open")

type CircuitBreaker struct {
	mu          sync.Mutex
	failures    int
	threshold   int
	openTimeout time.Duration
	openedAt    time.Time
}

func NewCircuitBreaker(threshold int, openTimeout time.Duration) *CircuitBreaker {
	if threshold < 1 {
		threshold = 5
	}
	if openTimeout <= 0 {
		openTimeout = 10 * time.Second
	}
	return &CircuitBreaker{threshold: threshold, openTimeout: openTimeout}
}

func (b *CircuitBreaker) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.failures < b.threshold {
		return true
	}
	if time.Since(b.openedAt) >= b.openTimeout {
		// Permit a trial request after the open interval.
		b.failures = b.threshold - 1
		return true
	}
	return false
}

func (b *CircuitBreaker) Success() {
	b.mu.Lock()
	b.failures = 0
	b.openedAt = time.Time{}
	b.mu.Unlock()
}

func (b *CircuitBreaker) Failure() {
	b.mu.Lock()
	b.failures++
	if b.failures >= b.threshold && b.openedAt.IsZero() {
		b.openedAt = time.Now()
	}
	b.mu.Unlock()
}

// Retry executes a bounded number of attempts. Callers should use it only for safe/idempotent operations.
func Retry(attempts int, delay time.Duration, operation func() error) error {
	if attempts < 1 {
		attempts = 1
	}
	var err error
	for i := 0; i < attempts; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		if i+1 < attempts {
			time.Sleep(delay)
		}
	}
	return err
}
