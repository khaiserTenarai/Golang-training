package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Secret hides its value whenever it's printed by accident.
type Secret string

func (s Secret) String() string { return "********" }
func (s Secret) Reveal() string { return string(s) }

// Config holds the ordinary, safe-to-print settings.
type Config struct {
	Env             string
	Port            string
	LogLevel        string
	ShutdownTimeout time.Duration
	FeatureBonus    bool // GET /employee/{id}/bonus
	FeatureSearch   bool // GET /employee?search=...
}

// Secrets is kept apart from Config so it never gets printed with it.
type Secrets struct {
	APIKey Secret
}

// readSetting looks at the environment first. If the variable isn't
// set, it asks the user, showing the default they get by pressing Enter.
func readSetting(reader *bufio.Reader, key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	fmt.Printf("%s is not set (Enter for %q): ", key, fallback)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return fallback
	}
	return text
}

func isTrue(value string) bool {
	switch strings.ToLower(value) {
	case "1", "true", "yes", "y", "on":
		return true
	}
	return false
}

// readSecret: API_KEY from the environment, then from the file named in
// API_KEY_FILE, and finally typed in by the user.
func readSecret(reader *bufio.Reader, key string) Secret {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return Secret(value)
	}
	if path, ok := os.LookupEnv(key + "_FILE"); ok && path != "" {
		if data, err := os.ReadFile(path); err == nil {
			return Secret(strings.TrimSpace(string(data)))
		}
	}

	fmt.Printf("%s is not set - type it in: ", key)
	text, _ := reader.ReadString('\n')
	return Secret(strings.TrimSpace(text))
}

func defaultLogLevel(env string) string {
	switch env {
	case "production":
		return "info"
	case "test":
		return "error"
	}
	return "debug"
}

// LoadConfig reads everything, applies the defaults and validates it.
func LoadConfig(reader *bufio.Reader) (Config, Secrets, []string) {
	env := readSetting(reader, "APP_ENV", "development")

	cfg := Config{
		Env:           env,
		Port:          readSetting(reader, "PORT", "8080"),
		LogLevel:      readSetting(reader, "LOG_LEVEL", defaultLogLevel(env)),
		FeatureBonus:  isTrue(readSetting(reader, "FEATURE_BONUS", "false")),
		FeatureSearch: isTrue(readSetting(reader, "FEATURE_SEARCH", "false")),
	}

	seconds, err := strconv.Atoi(readSetting(reader, "SHUTDOWN_TIMEOUT_SECONDS", "10"))
	if err != nil {
		seconds = 0 // validation below reports it
	}
	cfg.ShutdownTimeout = time.Duration(seconds) * time.Second

	secrets := Secrets{APIKey: readSecret(reader, "API_KEY")}

	return cfg, secrets, validate(cfg, secrets)
}

func validate(cfg Config, secrets Secrets) []string {
	var problems []string

	switch cfg.Env {
	case "development", "test", "production":
	default:
		problems = append(problems, fmt.Sprintf("APP_ENV must be development, test or production (got %q)", cfg.Env))
	}

	if port, err := strconv.Atoi(cfg.Port); err != nil || port < 1 || port > 65535 {
		problems = append(problems, "PORT must be a number between 1 and 65535")
	}

	switch cfg.LogLevel {
	case "debug", "info", "warn", "error":
	default:
		problems = append(problems, fmt.Sprintf("LOG_LEVEL must be debug, info, warn or error (got %q)", cfg.LogLevel))
	}

	if cfg.ShutdownTimeout <= 0 {
		problems = append(problems, "SHUTDOWN_TIMEOUT_SECONDS must be a whole number above 0")
	}

	minKeyLength := 8
	if cfg.Env == "production" {
		minKeyLength = 16
	}
	if len(secrets.APIKey.Reveal()) < minKeyLength {
		problems = append(problems, fmt.Sprintf("API_KEY must be at least %d characters in %s", minKeyLength, cfg.Env))
	}

	return problems
}
