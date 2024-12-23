package main

import (
	"go-jwt/configs"
	"go-jwt/models"
	"go-jwt/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix(`api/`).Subrouter()

	routes.AuthRoute(router)
	routes.UserRoute(router)
	routes.ProductRoute(router)
	routes.GetImage(router)

	configs.DB.AutoMigrate(&models.Product{})

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
