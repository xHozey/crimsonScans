package handlers

import (
	"encoding/json"
	"net/http"

	checker "github.com/xHozey/crimsonScans/funcs"
)

func Login(w http.ResponseWriter, r *http.Request) {
}

func (db *DataPass) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || db.Err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	if err := checker.EmailCheck(user.Email, db.Db); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if err := checker.PasswordCheck(user.Password); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if err := checker.UsernameCheck(user.Username, db.Db); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	
}
