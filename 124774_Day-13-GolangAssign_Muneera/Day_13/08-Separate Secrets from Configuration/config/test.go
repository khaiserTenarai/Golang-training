package config

// TestDefaults returns settings for automated testing.
func TestDefaults() *Config {
	return &Config{
		AppEnv:    "test",
		AppPort:   "8081",
		DBHost:    "localhost",
		DBPort:    "5432",
		DBUser:    "postgres",
		DBName:    "employee_test_db",
		DBSSLMode: "disable",
	}
}
