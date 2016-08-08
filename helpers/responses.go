package helpers

import (
	"encoding/json"
	"net/http"
)

// GetRes response of Get http handler
type GetRes struct {
	Value string `json:"value"`
}

// SetRes response of Set http handler
type SetRes struct {
	Key    string `json:"key"`
	Status string `json:"status"`
}

// TokenRes response of jwt http call
type TokenRes struct {
	Payload Token `json:"payload"`
}

// Response set the body and status code
func Response(body interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "Unable to create json body for response!", http.StatusBadRequest)
		return
	}
	// Set content type
	w.Header().Set("Content-Type", "application/json")
	// CORS Headerss
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if _, err := w.Write(json); err != nil {
		http.Error(w, "Unable to respond!", http.StatusInternalServerError)
	}
}
