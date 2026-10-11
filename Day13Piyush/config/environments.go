// Tasks 5, 6, 7: Environment-Specific Configurations
// Task 5 — Development configuration
// Task 6 — Test configuration
// Task 7 — Production configuration

package config

import "fmt"

func applyEnvironmentOverrides(cfg *Config) {
	switch cfg.Environment {
	case "development":
		applyDevDefaults(cfg)
	case "test":
		applyTestDefaults(cfg)
	case "production":
		applyProdDefaults(cfg)
	}
}

// Task 5: Development configuration
func applyDevDefaults(cfg *Config) {
	if cfg.LogLevel == "" {
		cfg.LogLevel = "debug" // Verbose logging in dev
	}
	if !cfg.EnableAuth {
		fmt.Println("  [DEV] Auth disabled — all endpoints open")
	}
	cfg.LogFormat = "text" // Human-readable in dev
}

// Task 6: Test configuration
func applyTestDefaults(cfg *Config) {
	cfg.LogLevel = "warn"   // Quiet during tests
	cfg.LogFormat = "text"
	cfg.EnableAuth = false  // No auth in tests
	cfg.MaxPageSize = 100
	cfg.ShutdownTimeoutSecs = 1 // Fast shutdown in tests
	fmt.Println("  [TEST] Test configuration applied")
}

// Task 7: Production configuration
func applyProdDefaults(cfg *Config) {
	cfg.LogLevel = "info"     // Only important logs
	cfg.LogFormat = "json"    // Structured for log aggregation
	cfg.EnableAuth = true     // Auth mandatory in prod
	cfg.MaxPageSize = 50
	cfg.ReadTimeoutSecs = 15
	cfg.WriteTimeoutSecs = 15
	cfg.ShutdownTimeoutSecs = 30 // Longer shutdown for draining
	fmt.Println("  [PROD] Production configuration applied")
}
