package routes

import (
	"github.com/gorilla/mux"
	"github.com/santosh/url-shortner/controllers"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/shortner", controllers.Shortner).Methods("POST")
	router.HandleFunc("/redirect/{url}", controllers.RedirectUrl).Methods("GET")
}
