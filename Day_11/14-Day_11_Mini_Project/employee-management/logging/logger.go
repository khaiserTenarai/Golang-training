// Package logging contains simple application logging functionality.
package logging

// Import the standard Go log package.
import "log"

// Logger wraps the standard Go logger.
type Logger struct {
	// logger stores the standard logger instance.
	logger *log.Logger
}

// New creates and returns a new application logger.
func New() *Logger {
	// Create a logger that includes date, time and source file information.
	return &Logger{
		// Initialize the standard logger.
		logger: log.Default(),
	}
}

// Info writes an informational message.
func (l *Logger) Info(message string) {
	// Print the informational message with an INFO prefix.
	l.logger.Printf("[INFO] %s", message)
}

// Error writes an error message.
func (l *Logger) Error(message string) {
	// Print the error message with an ERROR prefix.
	l.logger.Printf("[ERROR] %s", message)
}
