package routes

import (
	"go-jwt/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func ProductRoute(r *mux.Router) {
	router := r.PathPrefix("/api/product").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))).Subrouter()

	router.HandleFunc("/get-product", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/post-product", controllers.CreateProduct).Methods("POST")
}
