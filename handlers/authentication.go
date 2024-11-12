package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	checker "github.com/xHozey/crimsonScans/funcs"
	"golang.org/x/crypto/bcrypt"
)

func (db *DataPass) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || db.Err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	var username, password string
	err := db.Db.QueryRow("SELECT username, password FROM user WHERE email = ?", user.Email).Scan(&username, &password)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if !checker.PasswordValidation(password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := checker.GenerateJWT(username, db.Db)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	
	fmt.Fprint(w, token)
}

func (db *DataPass) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || db.Err != nil {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	if err := checker.EmailCheck(user.Email, db.Db); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if err := checker.PasswordCheck(user.Password); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if err := checker.UsernameCheck(user.Username, db.Db); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	stm, err := db.Db.Prepare("INSERT INTO user (username, email, password) VALUES(?,?,?)")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer stm.Close()

	_, err = stm.Exec(user.Username, user.Email, string(pass))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
