package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies = []Movie{
	{ID: "1", Isbn: "1234567890", Title: "The Shawshank Redemption", Director: &Director{FirstName: "Frank", LastName: "Darabont"}},
	{ID: "2", Isbn: "2345678901", Title: "The Godfather", Director: &Director{FirstName: "Francis", LastName: "Ford Coppola"}},
	{ID: "3", Isbn: "3456789012", Title: "Pulp Fiction", Director: &Director{FirstName: "Quentin", LastName: "Tarantino"}},
	{ID: "4", Isbn: "4567890123", Title: "The Dark Knight", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}},
	{ID: "5", Isbn: "5678901234", Title: "Fight Club", Director: &Director{FirstName: "David", LastName: "Fincher"}},
	{ID: "6", Isbn: "6789012345", Title: "Inception", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}},
	{ID: "7", Isbn: "7890123456", Title: "The Matrix", Director: &Director{FirstName: "Lana", LastName: "Wachowski"}},
	{ID: "8", Isbn: "8901234567", Title: "The Lord of the Rings: The Fellowship of the Ring", Director: &Director{FirstName: "Peter", LastName: "Jackson"}},
	{ID: "9", Isbn: "9012345678", Title: "Forrest Gump", Director: &Director{FirstName: "Robert", LastName: "Zemeckis"}},
	{ID: "10", Isbn: "0123456789", Title: "Schindler's List", Director: &Director{FirstName: "Steven", LastName: "Spielberg"}},
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")

	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	r.HandleFunc("/movies", createMovies).Methods("POST")

	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	r.HandleFunc("/movies/{id}", updateMovie).Methods("PATCH")

	fmt.Println("Server started on port 5000")

	log.Fatal(http.ListenAndServe(":5000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = params["id"]
	for index, m := range movies {
		if m.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
