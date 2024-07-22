package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santosh/forms/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	// Serve static files from the "assets" directory
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	routes.FormRoutes(r)

	http.Handle("/", r)

	fmt.Println("Go Running at : localhost:9000")

	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}
