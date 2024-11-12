package handlers

import "database/sql"

type DataPass struct {
	Db  *sql.DB
	Err error
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
