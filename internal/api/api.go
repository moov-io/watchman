package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ErrorResponse(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	return encodeError(w, err)
}

func encodeError(w http.ResponseWriter, err error) error {
	return json.NewEncoder(w).Encode(errorResponse{
		Error: err.Error(),
	})
}

func JsonResponse[T any](w http.ResponseWriter, object T) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(object)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		err = fmt.Errorf("problem rendering json: %w", err)

		fallbackErr := encodeError(w, err)
		if fallbackErr != nil {
			err = fmt.Errorf("problem rendering json marshal error: %w", fallbackErr)
		}

		return err
	}
	return nil
}
