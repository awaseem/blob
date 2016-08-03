package helpers

import (
	"encoding/json"
	"net/http"
)

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
