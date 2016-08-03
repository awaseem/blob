package main

import (
	"blob/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("Running on localhost:" + PORT)
	fmt.Println(http.ListenAndServe(":"+PORT, Handlers()))
}

// Handlers all route handlers for this service
func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/status", handlers.Status)
	return r
}
