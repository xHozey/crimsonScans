package handlers

import (
	"net/http"
)

func (db *DataPass) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if db.Err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("hello world"))
}
