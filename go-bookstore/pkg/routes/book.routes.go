
package routes;


import (
	"github.com/gorilla/mux"
	"example.com/go-bookstore/pkg/controllers"
)



var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST");
	router.HandleFunc("/book/",controllers.GetAllBooks).Methods("GET");
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE");
	router.HandleFunc("/book/{bookId}",controllers.UpdateBookById).Methods("PUT");
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET");
}