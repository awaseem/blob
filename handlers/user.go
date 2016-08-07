package handlers

import (
	"blob/helpers"
	"blob/jwtea"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// LoginUser login user for writes
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// parse incoming body from request
	var incomingBody helpers.User
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		http.Error(w, "Failed to parse request body!", http.StatusBadRequest)
		return
	}
	if errUnmarshal := json.Unmarshal(body, &incomingBody); errUnmarshal != nil {
		http.Error(w, "Failed to parse request body, incorrent format!", http.StatusBadRequest)
		return
	}
	// validate request parameters
	if incomingBody.Username == "" && incomingBody.Password == "" {
		http.Error(w, "Need a Username and Password", http.StatusBadRequest)
		return
	}
	// send success
	token, errToken := jwtea.Login(incomingBody.Username, incomingBody.Password)
	if errToken != nil {
		http.Error(w, "Failed to login!", http.StatusBadRequest)
		return
	}
	responseBody := helpers.Message{
		Success: true,
		Message: token,
	}
	helpers.Response(responseBody, w)
}
