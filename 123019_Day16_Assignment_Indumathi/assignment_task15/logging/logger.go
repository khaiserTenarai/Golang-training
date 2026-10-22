// Package logging provides application logging functionality.
package logging

import "log"

// Logger wraps standard Go logging.
type Logger struct {
	logger *log.Logger
}

// New creates a new Logger instance.
func New() *Logger {
	return &Logger{logger: log.Default()}
}

// Info logs informational messages with an [INFO] prefix.
func (l *Logger) Info(message string) {
	l.logger.Printf("[INFO] %s", message)
}

// Error logs error messages with an [ERROR] prefix.
func (l *Logger) Error(message string) {
	l.logger.Printf("[ERROR] %s", message)
}