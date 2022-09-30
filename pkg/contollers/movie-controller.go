package contollers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mertvasit/go-mongo-crud/pkg/models"
	"github.com/mertvasit/go-mongo-crud/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies, err := models.ReadMovies()
	if err != nil {
		log.Print("[ERROR]: Cannot fetch all movies - ", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot fetch all movies")
		return
	}
	res, _ := json.Marshal(movies)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie models.Movie
	utils.ParseRequestBody(r, &newMovie)

	res, err := newMovie.CreateMovie()
	if err != nil {
		log.Print("[ERROR]: Cannot insert movie - ", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot insert movie")
		return
	}

	resJson, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	id := mux.Vars(r)["id"]
	objectId, _ := primitive.ObjectIDFromHex(id)

	movie, err := models.ReadMovieById(objectId)
	if err != nil {
		log.Printf("[Error]: Cannot find movie "+id+" - ", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot find movie "+id)
		return
	}

	res, _ := json.Marshal(movie)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	objectId, _ := primitive.ObjectIDFromHex(id)

	var result models.Movie
	utils.ParseRequestBody(r, &result)

	mov, err := result.UpdateMovie(objectId)
	if err != nil {
		log.Print("[ERROR]: Cannot update movie "+id+" - ", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot update movie "+id)
		return
	}

	res, _ := json.Marshal(mov)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	objectId, _ := primitive.ObjectIDFromHex(id)

	mov, err := models.DeleteMovie(objectId)
	if err != nil {
		log.Print("[ERROR]: Cannot delete movie "+id+" - ", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot delete movie "+id)
		return
	}
	res, _ := json.Marshal(mov)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
