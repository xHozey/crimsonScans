package checker

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("ZvwWGK9YU6r8//f31libR9bWkwS2TSL3cn07uZBjbZw=")

func GenerateJWT(user string, db *sql.DB) (string, error) {
	var id int
	db.QueryRow("SELECT id FROM user WHERE username = ?", user).Scan(&id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":      id,
			"username": user,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}
	return username, nil
}
