package handlers

import (
	"blob/helpers"
	"blob/redis"
	"net/http"

	"github.com/gorilla/mux"
)

// Get http handler for redis gets
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	getRed, errRed := redis.Get(key)
	if errRed != nil {
		http.Error(w, "Failed to get value", http.StatusBadRequest)
		return
	}
	helpers.Response(helpers.GetRes{
		Value: getRed,
	}, w)
}
