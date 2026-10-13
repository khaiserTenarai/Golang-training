package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Secret string

func (s Secret) String() string { return "********" }

func (s Secret) Reveal() string { return string(s) }

type Config struct {
	DBHost string
	DBName string
	DBUser string
}

type Secrets struct {
	DBPassword Secret
	APIKey     Secret
}

func loadSecret(reader *bufio.Reader, key string) Secret {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return Secret(value)
	}

	if path, ok := os.LookupEnv(key + "_FILE"); ok && path != "" {
		data, err := os.ReadFile(path)
		if err == nil {
			return Secret(strings.TrimSpace(string(data)))
		}
		fmt.Printf("Couldn't read %s (%v)\n", path, err)
	}

	fmt.Printf("%s isn't set - type it in: ", key)
	text, _ := reader.ReadString('\n')
	return Secret(strings.TrimSpace(text))
}

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	cfg := Config{
		DBHost: ask(reader, "Database host: "),
		DBName: ask(reader, "Database name: "),
		DBUser: ask(reader, "Database user: "),
	}
	secrets := Secrets{
		DBPassword: loadSecret(reader, "DB_PASSWORD"),
		APIKey:     loadSecret(reader, "API_KEY"),
	}

	fmt.Printf("\nConfig:  %+v\n", cfg)
	fmt.Printf("Secrets: %+v\n", secrets)

	connString := fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.DBUser, secrets.DBPassword.Reveal(), cfg.DBHost, cfg.DBName)
	loggable := fmt.Sprintf("postgres://%s:%v@%s/%s", cfg.DBUser, secrets.DBPassword, cfg.DBHost, cfg.DBName)

	_ = connString // this is the one you'd hand to the database driver
	fmt.Println("\nConnecting with:", loggable)
}
