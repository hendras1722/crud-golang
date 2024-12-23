package routes

import (
	"go-jwt/controllers"
	"go-jwt/middleware"

	"github.com/gorilla/mux"
)

func ProductRoute(r *mux.Router) {
	router := r.PathPrefix("/api").Subrouter()
	router.Use(middleware.Auth)
	router.HandleFunc("/get-product", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/post-product", controllers.CreateProduct).Methods("POST")

}
