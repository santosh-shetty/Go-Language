package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santosh/book-store/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}
