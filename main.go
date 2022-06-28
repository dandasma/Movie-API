package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"log"
)

type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Year int `json:"year"`
	Director *Director `json:"director"`
	ImdbRating float64 `json:"imdb_rating"`
	}

type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

var movies []Movie

func main(){
	movies = append(movies ,Movie{
			ID: "1",
			Title: "The Shawshank Redemption",
			Year: 1994,
			Director: &Director{
				FirstName: "Frank",
				LastName: "Darabont",
			},
			ImdbRating: 9.3,
		},
		Movie{
			ID: "2",
			Title: "The Godfather",
			Year: 1972,
			Director: &Director{
				FirstName: "Francis",
				LastName: "Ford Coppola",
			},
			ImdbRating: 9.2,
		},
		Movie{
			ID: "3",
			Title: "Harry potter and the Philosopher's Stone",
			Year: 2001,
			Director: &Director{
				FirstName: "Chris",
				LastName: "Columbus",
			},
			ImdbRating: 8.7,
		},
	)

	router := mux.NewRouter()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//start the server
	log.Fatal(http.ListenAndServe(":8000", router))

	
}

//Generate a new id with the length of the movies slice
func getID() int {
	id := len(movies) + 1
	return id
}

//get all the movies
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

//get a movie by id
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	// sends an empty movie if the movie is not found with a 404 status code
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Movie not found")	
}

//create a new movie with a new id
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(getID())
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}



//update a movie by id with the new values from json body
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	// sends an empty movie if the movie is not found with a 404 status code
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Movie not found")
}

//delete a movie by id
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}
	// send an error message if the movie is not found with a 404 status code
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Movie not found")
}
