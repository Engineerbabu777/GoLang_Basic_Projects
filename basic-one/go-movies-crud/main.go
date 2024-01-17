

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
    firstname string `json:"firstname"`
	lastname string `json:"lastname"`
}

var movies []Movie;

func main(){

	router := mux.NewRouter();
	
	router.HandleFunc("/movies", getMovies).Methods("GET");
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET");
	router.HandleFunc("/movies", createMovie).Methods("POST");
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT");
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE");
	
	log.Fatal(http.ListenAndServe(":8000", router));
}