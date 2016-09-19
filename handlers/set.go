package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/awaseem/blob/helpers"
	"github.com/awaseem/blob/jwtea"
	"github.com/awaseem/blob/redis"
)

// Set http handler for redis set
func Set(w http.ResponseWriter, r *http.Request) {
	var incomingBody helpers.SetReq
	body, errIO := ioutil.ReadAll(r.Body)
	if errIO != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	if errUnmarshal := json.Unmarshal(body, &incomingBody); errUnmarshal != nil {
		http.Error(w, "Failed to prase body into json!", http.StatusBadRequest)
		return
	}
	if incomingBody.Token == "" || incomingBody.Key == "" || incomingBody.Value == "" {
		http.Error(w, "Missing values", http.StatusBadRequest)
		return
	}
	tokenCheck := jwtea.Check(incomingBody.Token)
	if !tokenCheck {
		http.Error(w, "Auth error!", http.StatusBadRequest)
		return
	}
	valRed, errRed := redis.Set(incomingBody.Key, incomingBody.Value)
	if errRed != nil {
		http.Error(w, "Failed to set value in redis", http.StatusBadRequest)
		return
	}
	helpers.Response(helpers.SetRes{
		Key:    incomingBody.Key,
		Status: valRed,
	}, w)
}
