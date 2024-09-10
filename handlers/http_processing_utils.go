package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/guilam34/financial_planner/models"
)

func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func encodeError(w http.ResponseWriter, status int, err error) {
	reqErr := models.RequestError{
		Error:   http.StatusText(status),
		Message: err.Error(),
	}
	encode(w, status, reqErr)
}
