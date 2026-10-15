package config

// ProductionDefaults returns settings for production.
// Actual database connection details must be provided during deployment.
func ProductionDefaults() *Config {
	return &Config{
		AppEnv:    "production",
		AppPort:   "8080",
		DBPort:    "5432",
		DBSSLMode: "require",
	}
}