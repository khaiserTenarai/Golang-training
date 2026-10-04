package main

import (
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

func main() {
	logger := Logger{
		Info:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		Warn:  log.New(os.Stdout, "WARN: ", log.LstdFlags),
		Error: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
	}

	logger.Info.Println("This is an info message")
	logger.Warn.Println("This is a warning message")
	logger.Error.Println("This is an error message")
}