package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"sync"
)

type rec struct {
	Name   string
	Salary float64
	Dept   string
}

func buggyTotal(emps []rec) float64 {
	sum := 0.0
	for i := 0; i <= len(emps); i++ {
		sum += emps[i].Salary
	}
	return sum
}

func buggyGroup(emps []rec) map[string]int {
	var counts map[string]int
	for _, e := range emps {
		counts[e.Dept]++
	}
	return counts
}

func buggyParse(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func total(emps []rec) float64 {
	sum := 0.0
	for i := 0; i < len(emps); i++ {
		sum += emps[i].Salary
	}
	return sum
}

func average(emps []rec) float64 {
	if len(emps) == 0 {
		return 0
	}
	return total(emps) / float64(len(emps))
}

func groupByDept(emps []rec) map[string]int {
	counts := make(map[string]int)
	for _, e := range emps {
		counts[e.Dept]++
	}
	return counts
}

func parseSalary(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid salary %q: %w", s, err)
	}
	return v, nil
}

func debugBad() {
	emps := []rec{{"Asha", 50000, "Eng"}, {"Ravi", 40000, "Ops"}}
	fmt.Println("parse abc ->", buggyParse("abc"))
	fmt.Println("total:", buggyTotal(emps))
	fmt.Println("group:", buggyGroup(emps))
}

func debugFixed() {
	emps := []rec{{"Asha", 50000, "Eng"}, {"Ravi", 40000, "Ops"}}
	fmt.Println("total:", total(emps))
	fmt.Println("avg:", average(emps))
	fmt.Println("avg empty:", average(nil))
	fmt.Println("groups:", groupByDept(emps))
	if _, err := parseSalary("abc"); err != nil {
		fmt.Println("error:", err)
	}
}

func logDemo() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("application started")

	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	textLogger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	jsonLogger := slog.New(slog.NewJSONHandler(f, opts))

	for _, l := range []*slog.Logger{textLogger, jsonLogger} {
		l.Debug("loading config", "path", "/etc/app.conf")
		l.Info("employee created", "id", 101, "name", "Asha")
		l.Warn("salary below threshold", "id", 101, "salary", 9000.50)
		l.Error("db failure", "err", errors.New("connection refused"), "retry", 3)
	}

	svc := jsonLogger.With("service", "employee-service", "version", "1.0")
	svc.Info("request handled",
		slog.Group("http", "method", "GET", "status", 200),
		slog.Int("duration_ms", 42),
	)
	fmt.Println("JSON logs written to app.log")
}

var counter int

func racyCounter() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println("racy counter (expected 1000):", counter)
}

func safeCounter() {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
		c  int
	)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			c++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("safe counter:", c)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . [debug-bad|debug|log|race-bad|race-safe]")
		return
	}
	switch os.Args[1] {
	case "debug-bad":
		debugBad()
	case "debug":
		debugFixed()
	case "log":
		logDemo()
	case "race-bad":
		racyCounter()
	case "race-safe":
		safeCounter()
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
