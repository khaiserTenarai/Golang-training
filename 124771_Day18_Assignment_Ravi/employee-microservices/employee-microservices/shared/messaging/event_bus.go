package messaging

import "log"

// Event is a tiny example of an asynchronous domain event. Task 9.
type Event struct {
	Type    string
	Payload any
}

// EventBus uses a Go channel to demonstrate asynchronous processing.
// Important: this is in-process only. Separate services need a broker such as RabbitMQ.
type EventBus struct{ events chan Event }

func NewEventBus(buffer int) *EventBus {
	if buffer < 1 {
		buffer = 1
	}
	return &EventBus{events: make(chan Event, buffer)}
}
func (b *EventBus) Publish(e Event) { b.events <- e }
func (b *EventBus) Consume(handler func(Event)) {
	go func() {
		for e := range b.events {
			handler(e)
		}
	}()
}
func ExampleNotificationConsumer() {
	bus := NewEventBus(5)
	bus.Consume(func(e Event) { log.Printf("notification consumer received event: %s", e.Type) })
	bus.Publish(Event{Type: "EmployeeCreated", Payload: map[string]string{"name": "Demo"}})
}
