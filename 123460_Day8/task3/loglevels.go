package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	flags := log.Ldate | log.Ltime | log.Lshortfile
	mw := io.MultiWriter(os.Stdout, file)

	// Distinct loggers for each level
	debug := log.New(mw, "DEBUG: ", flags)
	info := log.New(mw, "INFO:  ", flags)
	warn := log.New(mw, "WARN:  ", flags)
	errLog := log.New(mw, "ERROR: ", flags)

	debug.Println("Initializing configuration...")
	info.Println("Application started successfully.")
	warn.Println("High memory usage detected.")
	errLog.Println("Failed to connect to database.")
}