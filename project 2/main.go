package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(res http.ResponseWriter, req *http.Request) {

}

func getMovie(res http.ResponseWriter, req *http.Request) {

}

func createMovie(res http.ResponseWriter, req *http.Request) {

}

func updateMovie(res http.ResponseWriter, req *http.Request) {

}

func deleteMovie(res http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie {
		ID: "1",
		Isbn: "438227",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname: "Doe",
		},
	})

	movies = append(movies, Movie {
		ID: "2",
		Isbn: "45455",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Steve",
			Lastname: "Smith",
		},
	})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
