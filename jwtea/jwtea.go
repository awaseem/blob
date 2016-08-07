package jwtea

import (
	"blob/constants"
	"blob/helpers"
	"bytes"
	"encoding/json"
	"net/http"
)

const tokenURL string = "http://localhost:8000/create"
const header string = "application/json; charset=utf-8"

// Login create token based on a username and password
func Login(username string, password string) (string, error) {
	userCred := helpers.User{
		Username: username,
		Password: password,
	}
	b := new(bytes.Buffer)
	if errEncode := json.NewEncoder(b).Encode(userCred); errEncode != nil {
		return "", errEncode
	}
	res, errRes := http.Post(constants.GetJWTeaConfig()+"/login", header, b)
	if errRes != nil {
		return "", errRes
	}
	defer res.Body.Close()
	var tokenResponse helpers.TokenRes
	errToken := json.NewDecoder(res.Body).Decode(&tokenResponse)
	if errToken != nil {
		return "", errToken
	}
	return tokenResponse.Payload.Token, nil
}

// Check if token is valid return true otherwise false
func Check(token string) bool {
	userCred := helpers.Token{
		Token: token,
	}
	b := new(bytes.Buffer)
	if errEncode := json.NewEncoder(b).Encode(userCred); errEncode != nil {
		return false
	}
	res, errRes := http.Post(constants.GetJWTeaConfig()+"/check", header, b)
	if errRes != nil {
		return false
	}
	defer res.Body.Close()
	return true
}
