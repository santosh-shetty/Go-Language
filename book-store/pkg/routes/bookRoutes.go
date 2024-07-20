package routes

import (
	"github.com/gorilla/mux"
	"github.com/santosh/book-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.FindBookById).Methods("GET")
	router.HandleFunc("/book/create", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/delete/{id}", controllers.DeleteBookById).Methods("GET")
	router.HandleFunc("/book/update/{id}", controllers.UpdateBook).Methods("POST")
}
