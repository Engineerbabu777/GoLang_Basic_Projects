

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
   Isbn string  `json:"isbn"`
   Title string  `json:"title"`
   Director *Director `json:"director"`
}

type Director struct {
    Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie;

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(movies);
}

func getMovie(w http.ResponseWriter, r *http.Request){}
func createMovie(w http.ResponseWriter, r *http.Request){}
func updateMovie(w http.ResponseWriter, r *http.Request){}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json");
	params := mux.Vars(r);
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...);
		}
	}
}

func main(){

	router := mux.NewRouter();
	
	movies = append(movies, 
	Movie{ID: "1", Isbn: "438227", Title: "Ala Vaikuntapurmuloo", 
	Director: &Director{Firstname: "Trivikram ", Lastname: "Srinivas"}});
	// APPENDING / ADDING ANOTHER MOVIE TO THE MOVIES SLICES!
	movies = append(movies, 
	Movie{ID: "2", Isbn: "438287", Title: "Puspha", 
	Director: &Director{Firstname: "Sukumar ", Lastname: "Sukumar"}});


	router.HandleFunc("/movies", getMovies).Methods("GET");
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET");
	router.HandleFunc("/movies", createMovie).Methods("POST");
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT");
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE");
	
	fmt.Printf("Starting server at port 8000\n");
	log.Fatal(http.ListenAndServe(":8000", router));
}