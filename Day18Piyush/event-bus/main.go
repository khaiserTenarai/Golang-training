// Task 9: Design Asynchronous Communication
// Task 12: Distributed Transaction Handling
// Task 13: Saga Workflow
//
// Demonstrates an in-memory event bus with pub/sub pattern,
// saga orchestration, and compensating transactions.
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
// Task 9: Event Bus (Async Communication)
// ============================================================

type Event struct {
	Type      string
	Payload   map[string]interface{}
	Timestamp time.Time
}

type EventHandler func(Event)

type EventBus struct {
	mu       sync.RWMutex
	handlers map[string][]EventHandler
}

func NewEventBus() *EventBus {
	return &EventBus{handlers: make(map[string][]EventHandler)}
}

func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	event.Timestamp = time.Now()
	fmt.Printf("\n  [EventBus] Published: %s\n", event.Type)
	for _, handler := range eb.handlers[event.Type] {
		go handler(event) // Async execution
	}
}

// ============================================================
// Task 13: Saga Workflow — Employee Onboarding
// ============================================================

type SagaStep struct {
	Name       string
	Execute    func() error
	Compensate func() error // Rollback action
}

type Saga struct {
	Name          string
	Steps         []SagaStep
	completedSteps []int
}

func (s *Saga) Run() error {
	fmt.Printf("\n  [Saga] Starting: %s\n", s.Name)
	for i, step := range s.Steps {
		fmt.Printf("  [Saga] Step %d: %s — executing...\n", i+1, step.Name)
		if err := step.Execute(); err != nil {
			fmt.Printf("  [Saga] Step %d FAILED: %v\n", i+1, err)
			fmt.Println("  [Saga] Starting compensating transactions...")
			s.compensate()
			return fmt.Errorf("saga failed at step %d: %w", i+1, err)
		}
		s.completedSteps = append(s.completedSteps, i)
		fmt.Printf("  [Saga] Step %d: %s — done ✓\n", i+1, step.Name)
	}
	fmt.Printf("  [Saga] %s completed successfully ✓\n", s.Name)
	return nil
}

func (s *Saga) compensate() {
	// Rollback in reverse order
	for i := len(s.completedSteps) - 1; i >= 0; i-- {
		step := s.Steps[s.completedSteps[i]]
		fmt.Printf("  [Compensate] Rolling back: %s\n", step.Name)
		if err := step.Compensate(); err != nil {
			fmt.Printf("  [Compensate] WARNING: rollback failed for %s: %v\n", step.Name, err)
		}
	}
}

func main() {
	fmt.Println("===== Task 9: Asynchronous Event Bus =====")

	bus := NewEventBus()

	// Subscribe handlers (simulating different services listening)
	bus.Subscribe("employee.created", func(e Event) {
		fmt.Printf("  [Department-Service] Received %s: assigning desk for %v\n", e.Type, e.Payload["name"])
	})
	bus.Subscribe("employee.created", func(e Event) {
		fmt.Printf("  [Email-Service] Received %s: sending welcome email to %v\n", e.Type, e.Payload["email"])
	})
	bus.Subscribe("employee.created", func(e Event) {
		fmt.Printf("  [IT-Service] Received %s: provisioning laptop for %v\n", e.Type, e.Payload["name"])
	})
	bus.Subscribe("employee.deleted", func(e Event) {
		fmt.Printf("  [IT-Service] Received %s: revoking access for ID %v\n", e.Type, e.Payload["id"])
	})

	// Publish events
	bus.Publish(Event{Type: "employee.created", Payload: map[string]interface{}{"name": "Piyush", "email": "piyush@test.com"}})
	time.Sleep(200 * time.Millisecond)
	bus.Publish(Event{Type: "employee.deleted", Payload: map[string]interface{}{"id": 1}})
	time.Sleep(200 * time.Millisecond)

	// ============================================================
	// Task 13: Saga Workflow — Employee Onboarding (Success)
	// ============================================================
	fmt.Println("\n===== Task 13: Saga Workflow (Success) =====")

	successSaga := &Saga{
		Name: "Employee Onboarding",
		Steps: []SagaStep{
			{
				Name:       "Create Employee Record",
				Execute:    func() error { fmt.Println("    → Employee record created"); return nil },
				Compensate: func() error { fmt.Println("    ← Employee record deleted"); return nil },
			},
			{
				Name:       "Assign Department",
				Execute:    func() error { fmt.Println("    → Assigned to Engineering"); return nil },
				Compensate: func() error { fmt.Println("    ← Department assignment removed"); return nil },
			},
			{
				Name:       "Create Email Account",
				Execute:    func() error { fmt.Println("    → Email account created"); return nil },
				Compensate: func() error { fmt.Println("    ← Email account deactivated"); return nil },
			},
			{
				Name:       "Provision Equipment",
				Execute:    func() error { fmt.Println("    → Laptop provisioned"); return nil },
				Compensate: func() error { fmt.Println("    ← Equipment returned"); return nil },
			},
		},
	}
	successSaga.Run()

	// ============================================================
	// Task 12, 13: Saga with Failure + Compensation
	// ============================================================
	fmt.Println("\n===== Task 12: Distributed Transaction (Failure + Rollback) =====")

	failSaga := &Saga{
		Name: "Employee Transfer",
		Steps: []SagaStep{
			{
				Name:       "Debit Source Department Budget",
				Execute:    func() error { fmt.Println("    → Budget debited from Engineering"); return nil },
				Compensate: func() error { fmt.Println("    ← Budget credited back to Engineering"); return nil },
			},
			{
				Name:       "Update Employee Department",
				Execute:    func() error { fmt.Println("    → Employee moved to Marketing"); return nil },
				Compensate: func() error { fmt.Println("    ← Employee moved back to Engineering"); return nil },
			},
			{
				Name:       "Credit Target Department Budget",
				Execute:    func() error { return fmt.Errorf("Marketing budget is frozen") }, // FAILS
				Compensate: func() error { fmt.Println("    ← No credit to reverse"); return nil },
			},
		},
	}
	if err := failSaga.Run(); err != nil {
		fmt.Printf("\n  [Result] Transaction failed and rolled back: %v\n", err)
	}

	fmt.Println("\nDone!")
}
