package controllers

import (
	"encoding/json"
	"net/http"
	"restapi/pkg/models"
	"restapi/pkg/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var newBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	book := models.GetAllBooks()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookID := vars["bookId"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error, Invalid ID"))
		return
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	book := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookID := vars["bookId"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error, Invalid ID"))
		return
	}
	deletedBook := models.DeleteBook(ID)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	bookID := vars["bookId"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error, Invalid ID"))
		return
	}
	bookInfo, db := models.GetBookByID(ID)
	if updatedBook.Name != "" {
		bookInfo.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookInfo.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookInfo.Publication = updatedBook.Publication
	}
	db.Save(&bookInfo)
	res, _ := json.Marshal(bookInfo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Test(w http.ResponseWriter, r *http.Request) {

	for i := 60; i < 1000000; i++ {
		CreateBook := &models.Book{}
		utils.ParseBody(r, CreateBook)
		CreateBook.CreateBook()
		// res, _ := json.Marshal(book)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))
}
