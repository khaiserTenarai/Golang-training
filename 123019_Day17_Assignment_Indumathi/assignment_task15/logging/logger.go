package logging

import "log"

type Logger struct {
	logger *log.Logger
}

func New() *Logger {
	return &Logger{logger: log.Default()}
}

func (l *Logger) Info(msg string) {
	l.logger.Printf("[INFO] %s", msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Printf("[ERROR] %s", msg)
}