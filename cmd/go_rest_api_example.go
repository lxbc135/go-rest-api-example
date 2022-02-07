package main

import (
	"log"
	"net/http"

	"github.com/lxbc135/go_rest_api_example/handlers"

	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/articles", handlers.AllArticles)
	log.Fatal(http.ListenAndServe(":10000", r))
}

func main() {
	handleRequests()
}
