package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Prompter struct {
	reader *bufio.Reader
}

func NewPrompter() *Prompter {
	return &Prompter{reader: bufio.NewReader(os.Stdin)}
}

func (p *Prompter) Ask(label, fallback string) string {
	fmt.Printf("%s [%s]: ", label, fallback)
	text, _ := p.reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return fallback
	}
	return text
}

func (p *Prompter) AskBool(label string, fallback bool) bool {
	def := "n"
	if fallback {
		def = "y"
	}
	answer := strings.ToLower(p.Ask(label+" (y/n)", def))
	return answer == "y" || answer == "yes"
}

func fromEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func boolFromEnv(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return parsed
}
