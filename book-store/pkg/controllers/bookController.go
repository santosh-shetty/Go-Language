package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/santosh/book-store/pkg/models"
)

var BookStruct models.Book

// First Method for get
//
//	func GetBook(res http.ResponseWriter, req *http.Request) {
//		result := models.GetAllBook()
//		res.Header().Set("Content-Type", "application/json")
//		res.WriteHeader(http.StatusOK)
//		convertedRes, err := json.Marshal(result)
//		if err != nil {
//			http.Error(res, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		res.Write(convertedRes)
//	}
//

// Second Method for get
// GetBook retrieves all books.
func GetBook(res http.ResponseWriter, req *http.Request) {
	result := models.GetAllBook()
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(result)
}

// FindBookById retrieves a book by its ID.
func FindBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.ParseInt(vars["id"], 0, 0)
	result := models.GetBookById(id)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(result)
}

// CreateBook creates a new book.
func CreateBook(res http.ResponseWriter, req *http.Request) {
	var book models.Book
	if err := json.NewDecoder(req.Body).Decode(&book); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	createdBook := book.CreateBook()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(createdBook)
}

// DeleteBookById deletes a book by its ID.
func DeleteBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.ParseInt(vars["id"], 0, 0)

	if err := models.DeleteBook(id); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNoContent)
	json.NewEncoder(res).Encode(map[string]string{"message": "Deleted Successfully"})
}

// UpdateBook updates an existing book.
func UpdateBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var updatedBook models.Book
	if err := json.NewDecoder(req.Body).Decode(&updatedBook); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	book := models.GetBookById(id)
	if book.ID == 0 {
		http.Error(res, "Book not found", http.StatusNotFound)
		return
	}
	book.Name = updatedBook.Name
	book.Author = updatedBook.Author
	book.Publication = updatedBook.Publication
	book.UpdateBook()
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(book)
}
