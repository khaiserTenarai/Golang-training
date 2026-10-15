package config

// ProductionDefaults returns secure production-oriented defaults.
// Real database connection details must come from the environment.
func ProductionDefaults() *Config {
	return &Config{
		AppEnv:    "production",
		AppPort:   "8080",
		DBPort:    "5432",
		DBSSLMode: "require",
	}
}
