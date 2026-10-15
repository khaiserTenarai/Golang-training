package config

// TestDefaults returns isolated test configuration.
func TestDefaults() *Config {
	return &Config{
		AppEnv:     "test",
		AppPort:    "8081",
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBName:     "employee_test_db",
		DBSSLMode:  "disable",
	}
}
