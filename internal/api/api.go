package api

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ErrorResponse(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	return json.NewEncoder(w).Encode(errorResponse{
		Error: err.Error(),
	})
}

func JsonResponse[T any](w http.ResponseWriter, object T) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(object)
}
