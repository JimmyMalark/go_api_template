package httpx

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message" example:"internal server error"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, msg string, r *http.Request, err error) {
	slog.Warn(
		msg,
		"route", r.URL.Path,
		"status", status,
		"error", err,
	)

	_ = WriteJSON(w, status, ErrorResponse{
		Message: msg,
	})
}

type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}
