package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	port := ask(reader, "Enter the port (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}
	secs, err := strconv.Atoi(ask(reader, "Enter the shutdown timeout in seconds: "))
	if err != nil || secs <= 0 {
		fmt.Println("timeout must be a positive number")
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		fmt.Fprintln(w, "slow request finished")
	})
	srv := &http.Server{Addr: ":" + port, Handler: mux}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server error:", err)
			stop()
		}
	}()

	fmt.Printf("Server is running on port %s. Open /slow and press Ctrl+C\n", port)
	<-ctx.Done()

	fmt.Println("\nShutting down, waiting for running requests...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(secs)*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Println("shutdown did not finish in time:", err)
		return
	}
	fmt.Println("Server stopped gracefully")
}
