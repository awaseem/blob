package handlers

import (
	"fmt"
	"net/http"

	"github.com/awaseem/blob/helpers"
	"github.com/awaseem/blob/redis"

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
	helpers.Response(helpers.GetRes{Value: getRed}, w)
}

// GetSearch http handler for redis get search
func GetSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchKey := vars["key"]
	getRed, errRed := redis.SearchGet(searchKey)
	if errRed != nil {
		fmt.Println(errRed)
		http.Error(w, "Failed to get value", http.StatusBadRequest)
		return
	}
	helpers.Response(helpers.GetSearchRes{Value: getRed}, w)
}
