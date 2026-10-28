package eventbus

import (
    "log"
    "sync"
)

type Event struct {
    Name string `json:"name"`
    Data string `json:"data"`
}

type Handler func(Event)

type Bus struct {
    mu sync.RWMutex
    subscribers map[string][]Handler
    events chan Event
}

func New() *Bus {
    return &Bus{
        subscribers: make(map[string][]Handler),
        events: make(chan Event, 100),
    }
}

func (b *Bus) Subscribe(name string, handler Handler) {
    b.mu.Lock()
    defer b.mu.Unlock()
    b.subscribers[name] = append(b.subscribers[name], handler)
}

func (b *Bus) Publish(event Event) {
    // A bounded buffer provides back-pressure if consumers fall behind.
    b.events <- event
}

func (b *Bus) Start() {
    go func() {
        for event := range b.events {
            b.mu.RLock()
            handlers := append([]Handler(nil), b.subscribers[event.Name]...)
            b.mu.RUnlock()
            for _, handler := range handlers {
                h := handler
                go func() {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            log.Printf("event handler panic: %v", recovered)
                        }
                    }()
                    h(event)
                }()
            }
        }
    }()
}
