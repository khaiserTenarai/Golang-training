package main

import (
	"log"
	"log/slog"
	"os"
)

// Step 2: basic application logging
func basicLogging() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("application started")
	log.Printf("loaded %d employees", 42)
}

// Step 3: different log levels (standard library has none, so use prefixes)
func logLevels() {
	info := log.New(os.Stdout, "INFO:  ", log.Ldate|log.Ltime)
	warn := log.New(os.Stdout, "WARN:  ", log.Ldate|log.Ltime)
	errL := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	info.Println("service running")
	warn.Println("disk usage at 85%")
	errL.Println("failed to connect to database")
}

// Step 4: structured logs with log/slog (text and JSON)
func structuredLogs() {
	text := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	text.Debug("debugging", "step", 1)
	text.Info("employee created", "id", 101, "name", "Asha")
	text.Warn("salary above limit", "id", 101, "salary", 250000.0)
	text.Error("save failed", "id", 101, "err", "timeout")

	js := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	js.Info("employee created", "id", 102, "dept", "Engineering")
}

func main() {
	basicLogging()
	logLevels()
	structuredLogs()
}
