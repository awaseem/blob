package main

import (
	"blob/constants"
	"blob/handlers"
	"blob/redis"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("Running on localhost:" + constants.GetPortConfig())
	// create redis client
	redis.CreateClient()
	// start server
	fmt.Println(http.ListenAndServe(":"+constants.GetPortConfig(), Handlers()))
}

// Handlers all route handlers for this service
func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/status", handlers.Status)
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	return r
}
