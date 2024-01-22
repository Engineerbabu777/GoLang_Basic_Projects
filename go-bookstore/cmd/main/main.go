
package main;


import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"example.com/pkg/routes"
)



func main(){
	r := mux.NewRouter();
	routes.RegisterBookStoreRoutes(r);
	log.Println("Server starting at port 8080");
	http.Handle("/", r);

	log.Fatal(http.ListenAndServe("localhost:8080", r));
}