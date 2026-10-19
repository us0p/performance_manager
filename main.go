package main

import (
	"fmt"
	"html/template"
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

	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("%s: failed to load template.\n", err)
	}

	server := http.NewServeMux()

	server.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if tpl.Execute(res, nil) != nil {
			log.Println(err)
			http.Error(res, "Failed to excute template", http.StatusInternalServerError)
		}
	})

	fmt.Println("Starting server on port :8000")

	http.ListenAndServe(":8000", server)
}
