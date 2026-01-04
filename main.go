package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "forum_admin"
	password = "password123"
	dbname   = "forum_db"
)

func main() {
	// Connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initiate web server
	http.HandleFunc("/login", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // Wraps it in log to know so server crashes if somethings wrong
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Front end Formatting
	fmt.Fprintln(w, "Welcome to the Login Page!")

	// User Request handling
	requestType := r.Method
	if requestType == http.MethodPost {
		fmt.Println("Processing Login")
	} else {
		http.Error(w, "Method not allowed", 405)
	}
}
