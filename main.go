package main

import (
	"fmt"
	"log"
	"net/http"
	"performance_manager/src/database"
)

func main() {
	log.Println("Starting database connection...")
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("%s: failed to get database connection.\n", err)
	}
	log.Println("Database connected!")

	log.Println("Applying database migrations...")
	err = db.Migrate()
	if err != nil {
		log.Fatalf("%s: failed to apply migrations.\n", err)
	}
	log.Println("Migrations applied!")

	server := http.NewServeMux()

	server.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello, world!")
		return
	})

	fmt.Println("Starting server on port :8000")

	http.ListenAndServe(":8000", server)
}
