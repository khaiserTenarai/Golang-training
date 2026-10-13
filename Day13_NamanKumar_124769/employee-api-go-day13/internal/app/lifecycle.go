package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	HTTP            *http.Server
	ShutdownTimeout time.Duration
}

func (s *Server) Run() error {
	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("server_start address=%q", s.HTTP.Addr)
		serverErrors <- s.HTTP.ListenAndServe()
	}()

	signalContext, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-signalContext.Done():
		log.Printf("shutdown_signal_received")
	}

	shutdownContext, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()

	if err := s.HTTP.Shutdown(shutdownContext); err != nil {
		log.Printf("graceful_shutdown_failed error=%q", err.Error())
		if closeErr := s.HTTP.Close(); closeErr != nil {
			return closeErr
		}
		return err
	}
	log.Printf("shutdown_complete")
	return nil
}
