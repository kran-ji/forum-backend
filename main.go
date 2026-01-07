package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB // Declare database in global

type Credentials struct {
	Username string `json:"username"`
}

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

	var err error
	db, err = sql.Open("postgres", psqlInfo)
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
	// User Request handling
	requestType := r.Method
	if requestType != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request", 400)
		return
	}

	// Front end Formatting
	fmt.Fprintln(w, "Welcome to the Login Page!")
	fmt.Println(creds.Username)
}
