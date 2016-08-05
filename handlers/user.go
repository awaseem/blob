package handlers

import (
	"blob/helpers"
	"blob/redis"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser handler for creating a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// parse incoming body from request
	var incomingBody helpers.CreateUser
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to parse request body!", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &incomingBody); err != nil {
		http.Error(w, "Failed to parse request body, incorrent format!", http.StatusBadRequest)
		return
	}
	// validate request parameters
	if incomingBody.Username == "" && incomingBody.Password == "" {
		http.Error(w, "Need a Username and Password", http.StatusBadRequest)
		return
	}
	// generate password hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(incomingBody.Password), 10)
	if err != nil {
		http.Error(w, "Failed to create hash", http.StatusInternalServerError)
		return
	}
	// add username and password to database
	if success, err := redis.Set(incomingBody.Username, passwordHash); err != nil {
		http.Error(w, "Failed to set username to database", http.StatusInternalServerError)
		return
	} else if !success {
		http.Error(w, "Username already exists", http.StatusInternalServerError)
		return
	}
	// send success
	responseBody := helpers.Message{
		Success: true,
		Message: "Successfully created user!",
	}
	helpers.Response(responseBody, w)
}
