package main

import "fmt"

// These are our log levels.
// A bigger number means a more serious message.
const (
	DEBUG = 1
	INFO  = 2
	WARN  = 3
	ERROR = 4
)

// log prints a message if its level is allowed.
func log(currentLevel int, messageLevel int, message string) {

	// Show the message only when its level
	// is the same or higher than currentLevel.
	if messageLevel >= currentLevel {
		fmt.Println(message)
	}
}

func main() {

	// We want to show INFO, WARN, and ERROR.
	// DEBUG messages will not be shown.
	currentLevel := INFO

	// DEBUG is lower than INFO, so it will not print.
	log(currentLevel, DEBUG, "[DEBUG] Checking user")

	// INFO is allowed, so it will print.
	log(currentLevel, INFO, "[INFO] Server started")

	// WARN is allowed, so it will print.
	log(currentLevel, WARN, "[WARN] File is missing")

	// ERROR is allowed, so it will print.
	log(currentLevel, ERROR, "[ERROR] Database failed")
}
