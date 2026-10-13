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
	"syscall"
	"time"
)

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	port := ask(reader, "Port (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}
	slowSeconds, err := strconv.Atoi(ask(reader, "How many seconds should /slow take? "))
	if err != nil || slowSeconds < 0 {
		fmt.Println("Not a valid number, using 5.")
		slowSeconds = 5
	}
	waitSeconds, err := strconv.Atoi(ask(reader, "How many seconds may shutdown wait for active requests? "))
	if err != nil || waitSeconds <= 0 {
		fmt.Println("Not a valid number, using 10.")
		waitSeconds = 10
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello! Try /slow and then press Ctrl+C in the terminal.")
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  /slow request started")
		time.Sleep(time.Duration(slowSeconds) * time.Second)
		fmt.Fprintln(w, "Finished the slow request.")
		fmt.Println("  /slow request finished")
	})

	srv := &http.Server{Addr: ":" + port, Handler: mux}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server error:", err)
			stop()
		}
	}()
	fmt.Println("Server running on port", port, "- press Ctrl+C to stop it.")

	<-ctx.Done()
	fmt.Println("\nShutdown signal received. Waiting for active requests to finish...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(waitSeconds)*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Had to stop early:", err)
		return
	}
	fmt.Println("Server stopped cleanly.")
}
