package main

import "fmt"

// Task 4: Constructor-Based Dependency Injection
// Ensures structures are explicitly initialized with required references.
type Config struct {
	AppName string
}

type Application struct {
	cfg *Config
}

func NewApplication(cfg *Config) *Application {
	return &Application{cfg: cfg}
}

func (a *Application) Run() {
	fmt.Printf("[Task 4] App %s running with injected configuration.\n", a.cfg.AppName)
}

func main() {
	cfg := &Config{AppName: "Go-Architect-App"}
	app := NewApplication(cfg)
	app.Run()
}
