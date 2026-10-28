package httpclient

import (
	"errors"
	"sync"
	"time"
)

var ErrCircuitOpen = errors.New("circuit breaker is open")

type CircuitBreaker struct {
	mu               sync.Mutex
	failures         int
	threshold        int
	resetTimeout     time.Duration
	openedAt         time.Time
	open             bool
	halfOpenInFlight bool
}

func NewCircuitBreaker(threshold int, resetTimeout time.Duration) *CircuitBreaker {
	if threshold < 1 {
		threshold = 5
	}
	if resetTimeout <= 0 {
		resetTimeout = 30 * time.Second
	}

	return &CircuitBreaker{
		threshold:    threshold,
		resetTimeout: resetTimeout,
	}
}

func (b *CircuitBreaker) Allow() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.open {
		return nil
	}

	if time.Since(b.openedAt) < b.resetTimeout {
		return ErrCircuitOpen
	}

	if b.halfOpenInFlight {
		return ErrCircuitOpen
	}

	b.halfOpenInFlight = true
	return nil
}

func (b *CircuitBreaker) RecordSuccess() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.failures = 0
	b.open = false
	b.halfOpenInFlight = false
}

func (b *CircuitBreaker) RecordFailure() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.halfOpenInFlight = false
	b.failures++

	if b.failures >= b.threshold {
		b.open = true
		b.openedAt = time.Now()
	}
}
