package controller

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"employee-api/model"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeMessage(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// writeError translates errors from the service/repository into HTTP responses.
func writeError(w http.ResponseWriter, logger *slog.Logger, err error) {
	var ve *model.ValidationError
	switch {
	case errors.As(err, &ve):
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": ve.Error(), "fields": ve.Fields})
	case errors.Is(err, model.ErrNotFound):
		writeMessage(w, http.StatusNotFound, err.Error())
	case errors.Is(err, model.ErrConflict):
		writeMessage(w, http.StatusConflict, err.Error())
	case errors.Is(err, model.ErrFeatureDisabled):
		writeMessage(w, http.StatusNotFound, err.Error())
	case errors.Is(err, context.DeadlineExceeded):
		writeMessage(w, http.StatusGatewayTimeout, "the request took too long")
	case errors.Is(err, context.Canceled):
		writeMessage(w, http.StatusServiceUnavailable, "the request was cancelled")
	default:
		logger.Error("unexpected error", "error", err)
		writeMessage(w, http.StatusInternalServerError, "internal server error")
	}
}
