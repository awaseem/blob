package token

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type token struct {
	Token string `json:"token"`
}

type tokenBody struct {
	Username string `json:"username"`
}

type tokenRes struct {
	Payload token `json:"payload"`
}

const tokenURL string = "http://localhost:8000/create"
const header string = "application/json; charset=utf-8"

func CreateToken(username string) (string, error) {
	tokenBody := tokenBody{
		Username: username,
	}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(tokenBody)
	if err != nil {
		return "", err
	}
	res, errPost := http.Post(tokenURL, header, b)
	if errPost != nil {
		return "", err
	}
	defer res.Body.Close()
	var tokenResponse tokenRes
	errToken := json.NewDecoder(res.Body).Decode(&tokenResponse)
	if errToken != nil {
		return "", err
	}
	return tokenResponse.Payload.Token, nil
}
