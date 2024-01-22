

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

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	params := mux.Vars(r);
	for _, item := range movies{
       if(item.ID == params["id"]){
		json.NewEncoder(w).Encode(item);
		return
	   }
	}

}
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	var movie Movie;
	_ = json.NewDecoder(r.Body).Decode(&movie);
	movie.ID = strconv.Itoa(rand.Intn(10000000));
	movies = append(movies, movie);
	json.NewEncoder(w).Encode(movie);
}

// Define a function named updateMovie that handles an HTTP request for updating a movie.
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header of the response to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")

	// Get the parameters from the request using mux.Vars(r).
	params := mux.Vars(r)

	// Loop through the movies to find and update the movie with the specified ID.
	for index, item := range movies {
		// Check if the ID of the current movie matches the ID from the request parameters.
		if item.ID == params["id"] {
			// If a match is found, remove the movie from the movies slice.
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// Decode the JSON data from the request body into a new movie object.
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Set the ID of the movie to the ID from the request parameters.
	movie.ID = params["id"]

	// Append the updated movie to the movies slice.
	movies = append(movies, movie)

	// Encode the updated movie in JSON format and send it as the HTTP response.
	json.NewEncoder(w).Encode(movie)
}


func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json");
	params := mux.Vars(r);
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...);
			break;
		}
	}
	json.NewEncoder(w).Encode(movies);
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