package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Delete function
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.Id = strconv.Itoa(rand.Intn(10))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	/// creating some initial movies
	movie := &Movie{
		Id:    "1",
		Isbn:  "33885",
		Title: "Fresh",
		Director: &Director{
			Firstname: "John",
			Lastname:  "doe",
		},
	}
	movies = append(movies, *movie)
	r := mux.NewRouter()

	r.Path("/movies").HandlerFunc(getMovies).Methods("GET")
	r.Path("/movie/{id}").HandlerFunc(getMovie).Methods("GET")
	r.Path("/movies").HandlerFunc(createMovie).Methods("POST")
	r.Path("/movies/{id}").HandlerFunc(updateMovies).Methods("PUT")
	r.Path("/movies/{id}").HandlerFunc(deleteMovies).Methods("DELETE")

	fmt.Printf("Stsrting server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
