package config

import (
	"net"
	"strconv"
	"time"
)

const (
	EnvDevelopment = "development"
	EnvTest        = "test"
	EnvProduction  = "production"
)

// Config is the single, typed description of how the application runs.
// It is created once at startup, validated, and then passed down - nobody else calls os.Getenv.
type Config struct {
	App        AppConfig
	Server     ServerConfig
	Log        LogConfig
	Pagination PaginationConfig
	Features   Features
	Secrets    Secrets // kept apart from normal settings
}

type AppConfig struct {
	Name string
	Env  string
}

type ServerConfig struct {
	Host            string
	Port            int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	RequestTimeout  time.Duration // deadline given to every request's service/repository calls
	StartupTimeout  time.Duration // deadline for the startup initialization
	ShutdownTimeout time.Duration // how long in-flight requests may take after SIGINT/SIGTERM
}

type LogConfig struct {
	Level  string // debug | info | warn | error
	Format string // text | json
}

type PaginationConfig struct {
	DefaultPageSize int
	MaxPageSize     int
}

// Features are switches that change behaviour without changing code.
type Features struct {
	SearchEnabled bool // GET /employees?search=...
	ExposeSalary  bool // include "salary" in API responses
	SeedData      bool // insert sample employees at startup
}

// Secrets holds sensitive values only. Every field is a Secret so it can never be logged by accident.
type Secrets struct {
	APIKey Secret
}

func (c *Config) IsProduction() bool  { return c.App.Env == EnvProduction }
func (c *Config) IsDevelopment() bool { return c.App.Env == EnvDevelopment }
func (c *Config) IsTest() bool        { return c.App.Env == EnvTest }

func (s ServerConfig) Addr() string { return net.JoinHostPort(s.Host, strconv.Itoa(s.Port)) }

// Summary is a safe-to-print view of the configuration (secrets are redacted).
// It is used for the startup log and for GET /config.
func (c *Config) Summary() map[string]any {
	return map[string]any{
		"app": map[string]any{"name": c.App.Name, "env": c.App.Env},
		"server": map[string]any{
			"addr":            c.Server.Addr(),
			"readTimeout":     c.Server.ReadTimeout.String(),
			"writeTimeout":    c.Server.WriteTimeout.String(),
			"idleTimeout":     c.Server.IdleTimeout.String(),
			"requestTimeout":  c.Server.RequestTimeout.String(),
			"startupTimeout":  c.Server.StartupTimeout.String(),
			"shutdownTimeout": c.Server.ShutdownTimeout.String(),
		},
		"log":        map[string]any{"level": c.Log.Level, "format": c.Log.Format},
		"pagination": map[string]any{"defaultPageSize": c.Pagination.DefaultPageSize, "maxPageSize": c.Pagination.MaxPageSize},
		"features": map[string]any{
			"search":       c.Features.SearchEnabled,
			"exposeSalary": c.Features.ExposeSalary,
			"seedData":     c.Features.SeedData,
		},
		"secrets": map[string]any{"apiKey": c.Secrets.APIKey.String()},
	}
}
