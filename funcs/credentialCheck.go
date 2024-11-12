package checker

import (
	"database/sql"
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func PasswordCheck(pass string) error {
	if len(pass) < 8 || len(pass) > 36 {
		return errors.New("Bad Request")
	}
	return nil
}

func UsernameCheck(user string, db *sql.DB) error {
	exists := false
	db.QueryRow("SELECT 1 FROM user WHERE username = ?", user).Scan(&exists)
	if exists {
		return errors.New("Bad Request")
	}
	for _, val := range user {
		if (val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z') {
			continue
		} else {
			return errors.New("Bad Request")
		}
	}
	if len(user) < 3 || len(user) > 20 {
		return errors.New("Bad Request")
	}
	return nil
}

func EmailCheck(email string, db *sql.DB) error {
	exists := false
	db.QueryRow("SELECT 1 FROM user WHERE email = ?", email).Scan(&exists)
	if exists {
		return errors.New("Bad Request")
	}
	if !regexp.MustCompile(`^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$`).MatchString(email) || len(email) > 30 {
		return errors.New("Bad Request")
	}
	return nil
}

func PasswordValidation(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
