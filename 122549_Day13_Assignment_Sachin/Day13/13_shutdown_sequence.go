package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type resource struct {
	name string
	stop func(ctx context.Context) error
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Port (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	fmt.Print("Seconds allowed for the whole shutdown: ")
	waitText, _ := reader.ReadString('\n')
	waitSeconds, err := strconv.Atoi(strings.TrimSpace(waitText))
	if err != nil || waitSeconds <= 0 {
		fmt.Println("Not a valid number, using 10.")
		waitSeconds = 10
	}

	var resources []resource // in the order they were started

	// 1. database (pretend)
	fmt.Println("Starting: database")
	resources = append(resources, resource{"database", func(ctx context.Context) error {
		time.Sleep(300 * time.Millisecond)
		return nil
	}})

	// 2. background worker
	fmt.Println("Starting: background worker")
	workerCtx, cancelWorker := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-workerCtx.Done():
				return
			case <-ticker.C:
				fmt.Println("  worker: processing the queue...")
			}
		}
	}()
	resources = append(resources, resource{"background worker", func(ctx context.Context) error {
		cancelWorker()
		finished := make(chan struct{})
		go func() {
			wg.Wait()
			close(finished)
		}()
		select {
		case <-finished:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}})

	// 3. HTTP server
	fmt.Println("Starting: HTTP server on port", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running. Press Ctrl+C in the terminal to shut down.")
	})
	srv := &http.Server{Addr: ":" + port, Handler: mux}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server error:", err)
		}
	}()
	resources = append(resources, resource{"HTTP server", srv.Shutdown})

	signalCtx, stopSignals := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopSignals()
	<-signalCtx.Done()

	fmt.Println("\nShutting down in reverse order...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(waitSeconds)*time.Second)
	defer cancel()

	for i := len(resources) - 1; i >= 0; i-- {
		r := resources[i]
		fmt.Printf("  stopping %s... ", r.name)
		if err := r.stop(shutdownCtx); err != nil {
			fmt.Println("problem:", err)
		} else {
			fmt.Println("done")
		}
	}
	fmt.Println("Shutdown complete.")
}
