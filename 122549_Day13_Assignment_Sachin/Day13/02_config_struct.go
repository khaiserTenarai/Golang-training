package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	AppName        string
	Port           int
	Debug          bool
	LogLevel       string
	RequestTimeout time.Duration
}

// setting returns an environment variable if it's set. If it isn't,
// it asks the user to type the value in instead.
func setting(reader *bufio.Reader, key string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	fmt.Printf("%s is not set - type a value (or press Enter to leave it empty): ", key)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func loadConfig(reader *bufio.Reader) (Config, error) {
	appName := setting(reader, "APP_NAME")

	port, err := strconv.Atoi(setting(reader, "PORT"))
	if err != nil {
		return Config{}, fmt.Errorf("PORT has to be a whole number")
	}

	debug, err := strconv.ParseBool(setting(reader, "DEBUG"))
	if err != nil {
		return Config{}, fmt.Errorf("DEBUG has to be true or false")
	}

	logLevel := setting(reader, "LOG_LEVEL")

	seconds, err := strconv.Atoi(setting(reader, "REQUEST_TIMEOUT_SECONDS"))
	if err != nil {
		return Config{}, fmt.Errorf("REQUEST_TIMEOUT_SECONDS has to be a whole number")
	}

	return Config{
		AppName:        appName,
		Port:           port,
		Debug:          debug,
		LogLevel:       logLevel,
		RequestTimeout: time.Duration(seconds) * time.Second,
	}, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	cfg, err := loadConfig(reader)
	if err != nil {
		fmt.Println("Couldn't load the configuration:", err)
		return
	}

	fmt.Printf("\nLoaded configuration:\n%+v\n", cfg)
	fmt.Printf("\nThe app %q will listen on port %d.\n", cfg.AppName, cfg.Port)
}
