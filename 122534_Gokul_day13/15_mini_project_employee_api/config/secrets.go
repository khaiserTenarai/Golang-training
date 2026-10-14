package config

import (
	"fmt"
	"log/slog"
)

type Secrets struct {
	APIKey string
}

func LoadSecrets(p *Prompter, cfg Config) Secrets {
	if key := fromEnv("API_KEY", ""); key != "" {
		fmt.Println("API key read from the API_KEY environment variable")
		return Secrets{APIKey: key}
	}

	fallback := ""
	if cfg.Env != "production" {
		fallback = "dev-key-1234"
	}
	return Secrets{APIKey: p.Ask("API key", fallback)}
}

func (s Secrets) Validate(cfg Config) []string {
	var problems []string

	if s.APIKey == "" {
		problems = append(problems, "an API key is required")
	} else if cfg.Env == "production" && len(s.APIKey) < 8 {
		problems = append(problems, "the production API key must be at least 8 characters")
	}

	return problems
}

func (s Secrets) String() string {
	return "Secrets{APIKey: ****}"
}

func (s Secrets) LogValue() slog.Value {
	return slog.StringValue("****")
}
