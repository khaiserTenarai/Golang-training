// Task 13: Circuit Breaker — prevents cascading failures.
// States: CLOSED → OPEN → HALF-OPEN

package resilience

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type State int

const (
	Closed   State = iota
	Open
	HalfOpen
)

func (s State) String() string {
	return [...]string{"CLOSED", "OPEN", "HALF-OPEN"}[s]
}

type CircuitBreaker struct {
	mu               sync.Mutex
	name             string
	state            State
	failureCount     int
	failureThreshold int
	successCount     int
	successThreshold int
	lastFailure      time.Time
	cooldownPeriod   time.Duration
}

func NewCircuitBreaker(name string, failureThreshold int) *CircuitBreaker {
	return &CircuitBreaker{
		name: name, state: Closed,
		failureThreshold: failureThreshold, successThreshold: 2,
		cooldownPeriod: 10 * time.Second,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()
	switch cb.state {
	case Open:
		if time.Since(cb.lastFailure) > cb.cooldownPeriod {
			cb.state = HalfOpen
			cb.successCount = 0
			log.Printf("[CircuitBreaker:%s] OPEN → HALF-OPEN", cb.name)
			cb.mu.Unlock()
			return cb.tryCall(fn)
		}
		cb.mu.Unlock()
		return fmt.Errorf("[CircuitBreaker:%s] OPEN — request rejected", cb.name)
	case HalfOpen:
		cb.mu.Unlock()
		return cb.tryCall(fn)
	default:
		cb.mu.Unlock()
		return cb.tryCall(fn)
	}
}

func (cb *CircuitBreaker) tryCall(fn func() error) error {
	err := fn()
	cb.mu.Lock()
	defer cb.mu.Unlock()
	if err != nil {
		cb.failureCount++
		cb.lastFailure = time.Now()
		if cb.state == HalfOpen || cb.failureCount >= cb.failureThreshold {
			cb.state = Open
			log.Printf("[CircuitBreaker:%s] → OPEN (failures: %d)", cb.name, cb.failureCount)
		}
		return err
	}
	if cb.state == HalfOpen {
		cb.successCount++
		if cb.successCount >= cb.successThreshold {
			cb.state = Closed
			cb.failureCount = 0
			log.Printf("[CircuitBreaker:%s] HALF-OPEN → CLOSED", cb.name)
		}
	} else {
		cb.failureCount = 0
	}
	return nil
}

func (cb *CircuitBreaker) GetState() string {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state.String()
}
