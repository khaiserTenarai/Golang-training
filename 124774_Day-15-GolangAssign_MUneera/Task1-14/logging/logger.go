package logging

import "log"

type Logger struct {
	logger *log.Logger
}

func New() *Logger {
	return &Logger{
		logger: log.Default(),
	}
}

func (l *Logger) Info(message string) {
	l.logger.Printf("[INFO] %s", message)
}

func (l *Logger) Error(message string) {
	l.logger.Printf("[ERROR] %s", message)
}
