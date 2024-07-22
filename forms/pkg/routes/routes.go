package routes

import (
	"github.com/gorilla/mux"
	"github.com/santosh/forms/pkg/controllers"
	"github.com/santosh/forms/pkg/middleware"
)

func FormRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("GET", "POST")
	router.HandleFunc("/register", controllers.Register).Methods("GET", "POST")

	router.HandleFunc("/", controllers.Home).Methods("GET")
	protected := router.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/example", controllers.Example).Methods("GET")

}
