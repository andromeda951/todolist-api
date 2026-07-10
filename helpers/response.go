package helpers

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	body := map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	}

	_ = json.NewEncoder(w).Encode(body)
}

func Error(w http.ResponseWriter, statusCode int, message string, errors []string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	body := map[string]interface{}{
		"success": false,
		"message": message,
		"errors":  errors,
	}

	_ = json.NewEncoder(w).Encode(body)
}