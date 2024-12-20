package controllers

import (
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"
)

// Me returns the current user profile.
//
// The request context is expected to have "userinfo" key that contains
// the user information in the form of *helpers.MyCustomClaims.
//
// The response will be in JSON format and will contain the user's ID,
// name, and email.
func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	userResponse := &models.MyProfile{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	helpers.Response(w, 200, "My profile", userResponse)

}
