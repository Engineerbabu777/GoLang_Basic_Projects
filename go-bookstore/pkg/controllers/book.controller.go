
package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
	"example.com/pkg/utils"
	"example.com/pkg/models"
	"example.com/pkg/config"
);


var newBooks models.Book;


func GetAllBooks(w http.ResponseWriter, r *http.Request){
   newBooks	:= models.GetAllBooks();
   res,_ := json.Marshal(newBooks);
   w.Header().Set("Content-Type","pkglication/json");
   w.WriteHeader(http.StatusOK);
   w.Write(res);
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	bookId, err := strconv.ParseInt(params["bookId"],0,0);
    if err != nil {
		fmt.Println("Unable to convert the string into int");
		return;
	}
	bookDetails , _ := models.GetBookById(bookId);
	res, _ := json.Marshal(bookDetails);
	w.Header().Set("Content-Type","pkglication/json");
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{};
	utils.ParseBody(r,CreateBook);
	b := models.CreateBook(CreateBook);
	res, _ := json.Marshal(b);
	w.Header().Set("Content-Type","pkglication/json");
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}