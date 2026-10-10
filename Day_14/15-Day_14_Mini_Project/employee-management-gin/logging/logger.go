// Package logging provides application logging.
package logging

import "log"

// Logger wraps the standard logger.
type Logger struct{ logger *log.Logger }

// New creates a logger.
func New() *Logger { return &Logger{logger: log.Default()} }

// Info logs informational messages.
func (l *Logger) Info(message string) { l.logger.Printf("[INFO] %s", message) }

// Error logs error messages.
func (l *Logger) Error(message string) { l.logger.Printf("[ERROR] %s", message) }
