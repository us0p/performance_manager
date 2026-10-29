package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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

	// Should store all the pages of the application. Right now it wires the base template with each page body definition.
	pageTemplates := map[string]*template.Template{}
	pages, err := os.ReadDir("templates/pages")
	if err != nil {
		log.Fatalf("%s: failed to read template directory.\n", err)
	}

	for _, page := range pages {
		// ParseFiles will build one template set with allt the received files. Each file becomes one template inside the set, named after the file's name.
		// Each file is parsed in order, and the results go to the shared namespace which is a map from name to template.
		// the map is file's name -> set of named templates (defined by "define" or "block" constructs).
		// The template returned by ParseFiles points to the first file (index.html in this case).
		// In the code bellow, we're creating a set with:
		// - index.html: top-level text of the file.
		// - body: The body block in index.html which is overrided by the "body" block in the dynamic page.
		// - register.html: top-level text of the file.
		// - body: The defined "body" block in register.html file which overrides the "body" block in index.html.
		//
		// Name conflicts
		// - Only the file name is taken, not the path, so templates/index.html e templates/pages/index.html will both have "index.html" as the template name. Wichiever that's parsed last replaces the other silently.
		// - Same define block in multiple files. since the two files define the "body" block, the last one is wins.
		pageTemplates[page.Name()] = template.Must(template.ParseFiles(
			"templates/index.html",
			"templates/pages/"+page.Name(),
		))
	}

	server := http.NewServeMux()

	server.HandleFunc("/register", func(res http.ResponseWriter, req *http.Request) {
		// ExecuteTemplate allow us to select one of the named templates in the template set (in this case we have "index.html" and "register.html").
		// We're executing index.html because it'll render the top-level contents of this file and when it reaches the "body" block, it'll replace with the latest definition of that block in that template set.
		if err = pageTemplates["register.html"].ExecuteTemplate(res, "index.html", nil); err != nil {
			log.Println(err)
			http.Error(res, "Failed to excute template", http.StatusInternalServerError)
		}
	})

	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Starting server on port :8000")

	log.Fatal(http.ListenAndServe(":8000", server))
}
