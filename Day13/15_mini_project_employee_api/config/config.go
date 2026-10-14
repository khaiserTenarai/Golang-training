package config

import (
	"errors"
	"strconv"
	"strings"
)

type Config struct {
	Env             string
	Port            int
	LogLevel        string
	LogJSON         bool
	ShutdownTimeout int
	Features        Features
}

func presetFor(env string) Config {
	switch env {
	case "production":
		return Config{Env: env, Port: 8080, LogLevel: "info", LogJSON: true, ShutdownTimeout: 15}
	case "test":
		return Config{
			Env: env, Port: 8081, LogLevel: "error", ShutdownTimeout: 1,
			Features: Features{SalaryReport: true, DeleteEmployee: true},
		}
	default:
		return Config{
			Env: env, Port: 8080, LogLevel: "debug", ShutdownTimeout: 5,
			Features: Features{SalaryReport: true, DeleteEmployee: true},
		}
	}
}

func Load(p *Prompter) (Config, error) {
	env := strings.ToLower(p.Ask("Environment (development/test/production)", fromEnv("APP_ENV", "development")))
	cfg := presetFor(env)

	port, err := strconv.Atoi(p.Ask("Port", fromEnv("APP_PORT", strconv.Itoa(cfg.Port))))
	if err != nil {
		return cfg, errors.New("port must be a number")
	}
	cfg.Port = port

	cfg.LogLevel = strings.ToLower(p.Ask("Log level (debug/info/warn/error)", fromEnv("LOG_LEVEL", cfg.LogLevel)))

	secs, err := strconv.Atoi(p.Ask("Shutdown timeout in seconds", fromEnv("SHUTDOWN_TIMEOUT", strconv.Itoa(cfg.ShutdownTimeout))))
	if err != nil {
		return cfg, errors.New("shutdown timeout must be a number")
	}
	cfg.ShutdownTimeout = secs

	cfg.Features.SalaryReport = p.AskBool("Turn on the salary report feature", boolFromEnv("FEATURE_SALARY_REPORT", cfg.Features.SalaryReport))
	cfg.Features.DeleteEmployee = p.AskBool("Turn on the delete employee feature", boolFromEnv("FEATURE_DELETE_EMPLOYEE", cfg.Features.DeleteEmployee))

	return cfg, nil
}

func (c Config) Validate() []string {
	var problems []string

	switch c.Env {
	case "development", "test", "production":
	default:
		problems = append(problems, "environment must be development, test or production")
	}
	if c.Port < 1 || c.Port > 65535 {
		problems = append(problems, "port must be between 1 and 65535")
	}
	switch c.LogLevel {
	case "debug", "info", "warn", "error":
	default:
		problems = append(problems, "log level must be debug, info, warn or error")
	}
	if c.ShutdownTimeout < 1 {
		problems = append(problems, "shutdown timeout must be at least 1 second")
	}

	return problems
}
