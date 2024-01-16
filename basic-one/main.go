package main
import (
	"fmt"
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// IN CASE WE WILL CHECK!
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found!", http.StatusNotFound) 
		return
	}

	// CHECK FOR METHODS!
	if r.Method != "GET" {
		http.Error(w, "Method is not supported!", http.StatusNotFound)
		return;
	}

	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm();
	// IF ERROR OCCUR WHILE GETTING DATA FROM THE FORM!
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful");

	name := r.FormValue("name");
	address := r.FormValue("address");

	fmt.Fprintf(w, "Name: %s\n",name)
	fmt.Fprintf(w, "Address: %s\n",address)

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"));
	http.Handle("/", fileServer);
	http.HandleFunc("/form", formHandler);
	http.HandleFunc("/hello", helloHandler);

	fmt.Printf("Starting server for testing HTTP PORT 8080...\n")

	err := http.ListenAndServe(":8080", nil);

	if err != nil {
       log.Fatal(err);
	}
}