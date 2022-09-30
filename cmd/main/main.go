package main

import (
	"github.com/gorilla/mux"
	"github.com/mertvasit/go-mongo-crud/pkg/routes"
	"log"
	"net/http"
)

const PORT = ":8000"

func main() {
	router := mux.NewRouter()
	routes.MovieRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, router))
}
