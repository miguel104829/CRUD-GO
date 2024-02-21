package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to MySQL database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a simple HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Perform database operations
		// ...

		// Send response
		fmt.Fprintf(w, "Hello, miguel!")
	})

	// Start the server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
