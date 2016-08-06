package handlers

import (
	"blob/helpers"
	"blob/redis"
	"blob/token"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser handler for creating a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// parse incoming body from request
	var incomingBody helpers.User
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

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// parse incoming body from request
	var incomingBody helpers.User
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
	if result, err := redis.Get(incomingBody.Username); err != nil {
		http.Error(w, "Failed to to password from DB!", http.StatusInternalServerError)
	} else if result == "" {
		http.Error(w, "User does not exists!", http.StatusBadRequest)
	} else if errRedis := bcrypt.CompareHashAndPassword([]byte(result), []byte(incomingBody.Password)); errRedis != nil {
		http.Error(w, "User check failed!", http.StatusBadRequest)
	} else {
		// send success
		tokenStr, _ := token.CreateToken(incomingBody.Username)
		responseBody := helpers.Message{
			Success: true,
			Message: tokenStr,
		}
		helpers.Response(responseBody, w)
	}
}
