package main

import (
	"fmt"
	"net/http"

	"github.com/awaseem/blob/redis"
	"github.com/rs/cors"

	"github.com/awaseem/blob/handlers"

	"github.com/awaseem/blob/constants"

	"github.com/gorilla/mux"
)

func main() {
	println("Running on localhost:" + constants.GetPortConfig())
	// create redis client
	redis.CreateClient()
	// start server
	handler := cors.Default().Handler(Handlers())
	fmt.Println(http.ListenAndServe(":"+constants.GetPortConfig(), handler))
}

// Handlers all route handlers for this service
func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/status", handlers.Status)
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	r.HandleFunc("/get/{key}", handlers.Get).Methods("GET")
	r.HandleFunc("/get/search/{key}", handlers.GetSearch).Methods("GET")
	r.HandleFunc("/set", handlers.Set).Methods("POST")
	return r
}
