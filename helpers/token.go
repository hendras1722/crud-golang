package helpers

import (
	"fmt"
	"go-jwt/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var MySigningKey = []byte("myscretkey")

type MyCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {
	claims := MyCustomClaims{
		user.ID,
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(MySigningKey)
	return ss, err

}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	})

	claims, ok := token.Claims.(*MyCustomClaims)

	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	} else if !ok || !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
