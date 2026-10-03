package config
import (
	"os"
	"strings"
)

/*
	Config stores application configuration.
*/
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Day 9 concurrency configuration.
	Workers string
	Buffer  string
}

/*
	Load reads configuration from config.env.

	We are using file handling here.

	No bufio.
	No JSON.
	No external dotenv package.
*/
func Load() Config {

	/*
		Read the complete config.env file.
	*/
	data, _ := os.ReadFile("config.env")

	/*
		Convert file data into string
		and split it into lines.
	*/
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {

		/*
			Example:

			DB_HOST=localhost

			SplitN produces:

			DB_HOST
			localhost
		*/
		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			/*
				Store the value as an environment variable.
			*/
			os.Setenv(key, value)
		}
	}

	return Config{

		DBHost:     get("DB_HOST", "localhost"),

		DBPort: get("DB_PORT", "5432"),

		DBUser: get("DB_USER", "postgres"),

		DBPassword: get("DB_PASSWORD", "Root"),

		DBName: get("DB_NAME", "go_training_db"),

		DBSSLMode: get("DB_SSLMODE", "disable"),

		Workers: get("WORKERS", "3"),

		Buffer: get("BUFFER", "2"),
	}
}

/*
	DatabaseURL creates the PostgreSQL
	connection string.
*/
func (c Config) DatabaseURL() string {

	return "postgres://" +
		c.DBUser + ":" +
		c.DBPassword + "@" +
		c.DBHost + ":" +
		c.DBPort + "/" +
		c.DBName +
		"?sslmode=" + c.DBSSLMode
}

/*
	get reads an environment variable.

	If the value is empty,
	defaultValue is returned.
*/
func get(key string, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}