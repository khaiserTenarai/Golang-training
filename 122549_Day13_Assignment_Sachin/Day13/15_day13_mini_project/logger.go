package main

import (
	"log"
	"strings"
)

var levelOrder = map[string]int{"debug": 0, "info": 1, "warn": 2, "error": 3}

// Logger only prints messages at or above the configured level.
type Logger struct {
	level int
}

func NewLogger(level string) *Logger {
	return &Logger{level: levelOrder[level]}
}

func (l *Logger) logf(level, format string, args ...interface{}) {
	if levelOrder[level] < l.level {
		return
	}
	log.Printf("[%s] "+format, append([]interface{}{strings.ToUpper(level)}, args...)...)
}

func (l *Logger) Debugf(format string, args ...interface{}) { l.logf("debug", format, args...) }
func (l *Logger) Infof(format string, args ...interface{})  { l.logf("info", format, args...) }
func (l *Logger) Warnf(format string, args ...interface{})  { l.logf("warn", format, args...) }
func (l *Logger) Errorf(format string, args ...interface{}) { l.logf("error", format, args...) }
