package routes

import (
	"github.com/gorilla/mux"
	"github.com/mertvasit/go-mongo-crud/pkg/contollers"
)

var MovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie", contollers.GetMovies).Methods("GET")
	router.HandleFunc("/movie", contollers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", contollers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{id}", contollers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{id}", contollers.DeleteMovie).Methods("DELETE")
}
