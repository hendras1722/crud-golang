package routes

import (
	"go-jwt/controllers"

	"github.com/gorilla/mux"
)

func GetImage(r *mux.Router) {
	r.PathPrefix("/uploads/").Handler(controllers.StaticFileHandler())
}
