package config

// ProductionDefaults returns production-oriented settings.
// Real database connection details must be provided at deployment.
func ProductionDefaults() *Config {
	return &Config{
		AppEnv:    "production",
		AppPort:   "8080",
		DBPort:    "5432",
		DBSSLMode: "require",
	}
}
