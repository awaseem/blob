package handlers

import (
	"blob/helpers"
	"net/http"
)

// Status send the status of the service
func Status(w http.ResponseWriter, r *http.Request) {
	status := helpers.Message{
		Success: true,
		Message: "OK",
	}
	helpers.Response(status, w)
}
