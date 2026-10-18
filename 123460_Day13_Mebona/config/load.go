package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// LookupFunc finds a setting by name. os.LookupEnv has exactly this shape,
// and tests pass a map-backed one, so configuration is easy to test.
type LookupFunc func(key string) (string, bool)

// Load reads the configuration from the current directory:
// real environment > .env.secrets > .env.<APP_ENV> > .env > built-in defaults.
func Load() (*Config, error) { return LoadFromDir(".") }

func LoadFromDir(dir string) (*Config, error) {
	base, err := readDotEnv(filepath.Join(dir, ".env"))
	if err != nil {
		return nil, err
	}
	secrets, err := readDotEnv(filepath.Join(dir, ".env.secrets"))
	if err != nil {
		return nil, err
	}

	// Which environment? Real env var first, then .env, then development.
	env := strings.ToLower(strings.TrimSpace(os.Getenv("APP_ENV")))
	if env == "" {
		env = strings.ToLower(strings.TrimSpace(base["APP_ENV"]))
	}
	if env == "" {
		env = EnvDevelopment
	}

	overlay := map[string]string{}
	if env == EnvDevelopment || env == EnvTest || env == EnvProduction { // never build a path from unknown input
		if overlay, err = readDotEnv(filepath.Join(dir, ".env."+env)); err != nil {
			return nil, err
		}
	}

	layers := []map[string]string{secrets, overlay, base}
	lookup := func(key string) (string, bool) {
		if v, ok := os.LookupEnv(key); ok {
			return v, true
		}
		for _, m := range layers {
			if v, ok := m[key]; ok {
				return v, true
			}
		}
		return "", false
	}
	return LoadFrom(lookup)
}

// LoadFrom builds and validates a Config from any lookup source.
func LoadFrom(lookup LookupFunc) (*Config, error) {
	l := &loader{lookup: lookup}

	env := strings.ToLower(l.str("APP_ENV", EnvDevelopment))
	cfg := defaultsFor(env) // implementation of "configuration defaults" per environment
	cfg.App.Env = env

	cfg.App.Name = l.str("APP_NAME", cfg.App.Name)

	cfg.Server.Host = l.str("SERVER_HOST", cfg.Server.Host)
	cfg.Server.Port = l.integer("SERVER_PORT", cfg.Server.Port)
	cfg.Server.ReadTimeout = l.duration("SERVER_READ_TIMEOUT", cfg.Server.ReadTimeout)
	cfg.Server.WriteTimeout = l.duration("SERVER_WRITE_TIMEOUT", cfg.Server.WriteTimeout)
	cfg.Server.IdleTimeout = l.duration("SERVER_IDLE_TIMEOUT", cfg.Server.IdleTimeout)
	cfg.Server.RequestTimeout = l.duration("REQUEST_TIMEOUT", cfg.Server.RequestTimeout)
	cfg.Server.StartupTimeout = l.duration("STARTUP_TIMEOUT", cfg.Server.StartupTimeout)
	cfg.Server.ShutdownTimeout = l.duration("SHUTDOWN_TIMEOUT", cfg.Server.ShutdownTimeout)

	cfg.Log.Level = strings.ToLower(l.str("LOG_LEVEL", cfg.Log.Level))
	cfg.Log.Format = strings.ToLower(l.str("LOG_FORMAT", cfg.Log.Format))

	cfg.Pagination.DefaultPageSize = l.integer("PAGE_SIZE_DEFAULT", cfg.Pagination.DefaultPageSize)
	cfg.Pagination.MaxPageSize = l.integer("PAGE_SIZE_MAX", cfg.Pagination.MaxPageSize)

	cfg.Features.SearchEnabled = l.boolean("FEATURE_SEARCH", cfg.Features.SearchEnabled)
	cfg.Features.ExposeSalary = l.boolean("FEATURE_EXPOSE_SALARY", cfg.Features.ExposeSalary)
	cfg.Features.SeedData = l.boolean("FEATURE_SEED_DATA", cfg.Features.SeedData)

	cfg.Secrets.APIKey = Secret(l.str("API_KEY", cfg.Secrets.APIKey.Reveal()))

	if len(l.errs) > 0 { // unparsable values (e.g. SERVER_PORT=abc)
		return nil, errors.Join(l.errs...)
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// loader reads typed values and remembers every problem, so the user sees ALL mistakes at once.
type loader struct {
	lookup LookupFunc
	errs   []error
}

func (l *loader) raw(key string) (string, bool) {
	v, ok := l.lookup(key)
	v = strings.TrimSpace(v)
	return v, ok && v != "" // an empty value counts as "not set"
}

func (l *loader) str(key, def string) string {
	if v, ok := l.raw(key); ok {
		return v
	}
	return def
}

func (l *loader) integer(key string, def int) int {
	v, ok := l.raw(key)
	if !ok {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		l.errs = append(l.errs, fmt.Errorf("%s: %q is not a valid integer", key, v))
		return def
	}
	return n
}

func (l *loader) boolean(key string, def bool) bool {
	v, ok := l.raw(key)
	if !ok {
		return def
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		l.errs = append(l.errs, fmt.Errorf("%s: %q is not a valid boolean (use true or false)", key, v))
		return def
	}
	return b
}

func (l *loader) duration(key string, def time.Duration) time.Duration {
	v, ok := l.raw(key)
	if !ok {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		l.errs = append(l.errs, fmt.Errorf("%s: %q is not a valid duration (examples: 500ms, 5s, 2m)", key, v))
		return def
	}
	return d
}
