package logging

import (
	"log"
)

type Logger interface {
	Info(v ...interface{})
	Error(v ...interface{})
}

type stdLogger struct{}

func New() Logger {
	return &stdLogger{}
}

func (l *stdLogger) Info(v ...interface{}) {
	log.Println(append([]interface{}{"[INFO]"}, v...)...)
}

func (l *stdLogger) Error(v ...interface{}) {
	log.Println(append([]interface{}{"[ERROR]"}, v...)...)
}