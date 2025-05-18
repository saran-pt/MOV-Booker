package main

import (
	"fmt"
	"log"
	"net/http"
	// "strconv"
	"encoding/json"
	// "reflect"
	"github.com/gorilla/mux"
)

type Movie struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Movie Buzz...");
}


func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}


func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, obj := range movies {
		if obj.Id == params["id"]{
			json.NewEncoder(w).Encode(obj);
			return
		}
	}
	fmt.Fprintf(w, "No Data Found")
}


func createMovie(w http.ResponseWriter, r *http.Request) {
	params := http.NewDecoder(r).Decode(&Movie)
	fmt.Println(params)
}


var movies []Movie

func main() {
	// Movie Data
	movies = append(movies, Movie{"101", "OkOK", &Director{"Ram", "Gopal"}})
	movies = append(movies, Movie{"102", "Mind.Off", &Director{"Pream", "Hank"}})
	fmt.Println("Server is running...");

	r := mux.NewRouter()
	
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies", createMovie).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r));
}
