package models

import (
	"example.com/pkg/config"
	"github.com/jinzhu/gorm"
)


var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publications string `json:"publications"`
}


func InitDB() {
	config.ConnectDB();
	db := config.GetDB();
	db.AutoMigrate(&Book{});
}

func CreateBook(b *Book) *Book{
    db.NewRecord(b);
	db.Create(&b);
	return b;
}

func GetAllBooks() Book[]{
	var books []Book;
	db.Find(&books);
	return books;
}

func GetBookById(id int64, ) (*Book, *gorm.DB){
	var book Book;
	db := db.Where("ID=?", id).Find(&book);
	return book,db;
}

func UpdateBook(book *Book, id int64) Book{
	db.Save(&book);
	return *book;
}

func DeleteBook(id int64) Book{
	var book Book;
	db.Where("ID=?", id).Delete(book);
	return book;
}