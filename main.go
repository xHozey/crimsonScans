package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/xHozey/crimsonScans/db"
	"github.com/xHozey/crimsonScans/handlers"
)

func main() {
	db, err := database.OpenDB()
	database := handlers.DataPass{Db: db, Err: err}

	mux := http.NewServeMux()
	serverConfig := &http.Server{Addr: ":8080", Handler: mux}

	mux.HandleFunc("/", database.HomeHandler)
	mux.HandleFunc("/api", handlers.DataDisplay)
	mux.HandleFunc("/login", database.Login)
	mux.HandleFunc("/register", database.Register)

	fmt.Println("Serve is listening on port 8080")
	log.Fatal(serverConfig.ListenAndServe())
}
