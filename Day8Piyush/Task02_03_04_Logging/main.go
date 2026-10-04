// Tasks 2, 3, 4: Application Logging, Log Levels, and Structured Logs
// Task 2 — Basic application logging using the standard log package
// Task 3 — Implementing different log levels (DEBUG, INFO, WARN, ERROR)
// Task 4 — Creating structured logs in JSON format
// Run: go run main.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// ============================================================
// Task 2: Basic Application Logging
// ============================================================

func demoBasicLogging() {
	fmt.Println("========== TASK 2: Basic Application Logging ==========")

	// Default logger writes to stderr
	log.Println("This is a basic log message")
	log.Printf("Employee %s has been created with ID %d", "Piyush", 101)

	// Log to a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Cannot open log file:", err)
	}
	defer file.Close()

	// Write to both file and console
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Logging to both console and file")
	log.Println("Application started successfully")

	// Reset to default
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags)

	fmt.Println()
}

// ============================================================
// Task 3: Different Log Levels
// ============================================================

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type LevelLogger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	minLevel    LogLevel
}

func NewLevelLogger(minLevel LogLevel) *LevelLogger {
	return &LevelLogger{
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:  log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger:  log.New(os.Stdout, "[WARN]  ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		minLevel:    minLevel,
	}
}

func (l *LevelLogger) Debug(msg string) {
	if l.minLevel <= DEBUG {
		l.debugLogger.Println(msg)
	}
}

func (l *LevelLogger) Info(msg string) {
	if l.minLevel <= INFO {
		l.infoLogger.Println(msg)
	}
}

func (l *LevelLogger) Warn(msg string) {
	if l.minLevel <= WARN {
		l.warnLogger.Println(msg)
	}
}

func (l *LevelLogger) Error(msg string) {
	if l.minLevel <= ERROR {
		l.errorLogger.Println(msg)
	}
}

func demoLogLevels() {
	fmt.Println("========== TASK 3: Different Log Levels ==========")

	// Show all levels (DEBUG and above)
	fmt.Println("--- Min level: DEBUG (shows everything) ---")
	logger := NewLevelLogger(DEBUG)
	logger.Debug("Fetching employee with ID 101")
	logger.Info("Employee Piyush loaded successfully")
	logger.Warn("Salary is below minimum threshold")
	logger.Error("Failed to update employee record")

	// Show only WARN and above
	fmt.Println("\n--- Min level: WARN (only WARN and ERROR) ---")
	logger2 := NewLevelLogger(WARN)
	logger2.Debug("This DEBUG message will NOT appear")
	logger2.Info("This INFO message will NOT appear")
	logger2.Warn("This WARN message WILL appear")
	logger2.Error("This ERROR message WILL appear")

	fmt.Println()
}

// ============================================================
// Task 4: Structured Logs (JSON format)
// ============================================================

type StructuredLog struct {
	Timestamp string      `json:"timestamp"`
	Level     string      `json:"level"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Caller    string      `json:"caller,omitempty"`
}

type StructuredLogger struct{}

func (sl *StructuredLogger) log(level, message, caller string, data interface{}) {
	entry := StructuredLog{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		Data:      data,
		Caller:    caller,
	}
	jsonBytes, _ := json.Marshal(entry)
	fmt.Println(string(jsonBytes))
}

func (sl *StructuredLogger) Info(msg string, data interface{}) {
	sl.log("INFO", msg, "main.go", data)
}

func (sl *StructuredLogger) Error(msg string, data interface{}) {
	sl.log("ERROR", msg, "main.go", data)
}

func (sl *StructuredLogger) Warn(msg string, data interface{}) {
	sl.log("WARN", msg, "main.go", data)
}

func (sl *StructuredLogger) Debug(msg string, data interface{}) {
	sl.log("DEBUG", msg, "main.go", data)
}

func demoStructuredLogging() {
	fmt.Println("========== TASK 4: Structured Logs (JSON) ==========")

	slog := &StructuredLogger{}

	slog.Info("Application started", map[string]string{"version": "1.0.0", "env": "production"})

	slog.Info("Employee created", map[string]interface{}{
		"employee_id": 101,
		"name":        "Piyush",
		"department":  "Engineering",
	})

	slog.Warn("High memory usage", map[string]interface{}{
		"usage_percent": 85.5,
		"threshold":     80.0,
	})

	slog.Error("Database connection failed", map[string]interface{}{
		"host":    "localhost",
		"port":    5432,
		"retries": 3,
	})

	slog.Debug("Query executed", map[string]interface{}{
		"query":    "SELECT * FROM employees",
		"duration": "45ms",
		"rows":     150,
	})

	fmt.Println()
}

// ============================================================
// Main
// ============================================================

func main() {
	demoBasicLogging()
	demoLogLevels()
	demoStructuredLogging()

	fmt.Println("All logging tasks completed! Check app.log for file output.")
}
