package router

import (
	"github.com/aniketDinda/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	//Routing
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteAllMovies", controller.DeleteMovies).Methods("DELETE")

	return r
}
