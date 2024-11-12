package handlers

import (
	"fmt"
	"net/http"

	checker "github.com/xHozey/crimsonScans/funcs"
)

func (db *DataPass) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if db.Err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]
	username, err := checker.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}
	w.Write([]byte("hello there " + username))
}
