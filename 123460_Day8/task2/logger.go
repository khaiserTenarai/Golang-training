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

	// Log to both console and file with timestamps and file names
	flags := log.Ldate | log.Ltime | log.Lshortfile
	mw := io.MultiWriter(os.Stdout, file)

	info := log.New(mw, "INFO:  ", flags)
	warn := log.New(mw, "WARN:  ", flags)

	info.Println("Application started successfully.")
	info.Printf("User %s logged in from IP %s", "admin", "192.168.1.50")
	warn.Println("High memory usage detected.")
}