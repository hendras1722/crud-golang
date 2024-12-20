package routes

import (
	"go-jwt/controllers"
	"go-jwt/middleware"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	router := r.PathPrefix("/api/user").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("/me", controllers.Me).Methods("GET")
}
