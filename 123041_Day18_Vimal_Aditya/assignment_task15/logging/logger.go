package logging

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string, v ...interface{}) {
	l.infoLogger.Printf(msg, v...)
}

func (l *Logger) Error(msg string, v ...interface{}) {
	l.errorLogger.Printf(msg, v...)
}