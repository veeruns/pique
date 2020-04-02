package pique

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func startserver() {
	router := mux.NewRouter()

	router.HandleFunc("/", index)
	router.HandleFunc("/posts", getPosts)
	logwriter := log.Writer()
	loggedRouter := handlers.LoggingHandler(logwriter, router)
	http.ListenAndServe("localhost:8080", loggedRouter)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting to the posts  ")
	fmt.Fprintf(w, "Hello World from the http server\n")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from pique index page")
}
