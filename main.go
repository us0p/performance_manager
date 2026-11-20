package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"performance_manager/src/api"
	brules "performance_manager/src/bRules"
	"performance_manager/src/controllers"
	"performance_manager/src/database"
	"performance_manager/src/database/repositories"
)

// main.go is the composition root

// Compostion root is the single place where all the concrete objects are created and connected to each other.
// Since inner circles depend only on abstractions, there must be a place that know the concrete type and actually wire them together.
// The composition root is the only component allowed to know every concrete type, so all the other components don't have to.
// The composition root is usually messy and convoluted but it enables:
//   - Testability: a test can compose the same use cases with fakes, effectively acting as a second composition root.
//   - Replaceability: swapping Postgres or net/http means changing the wiring in one file.
//   - Visibility: the whole dependency graph can be read in one place, from top to bottom.
//   - Pluggability: in Bob's words, main is a plugin to the application. You can have one main for production, another for dev, and another for a CLI,
//     all reusing the same inner layers.
//
// Rules of thumb and anti-patterns:
// - Keep only one, as close to the entry point as possible. Building objects in many places scatters your wiring decisions.
// - Build the graph once at startup. Don't create repositories per request inside handlers.
// - Don't let use cases construct their own dependencies.
// - Avoid a service locator. Don't pass the whole App or a container around so that code can pull whatever it needs. Pass each component only the
// dependencies it actually uses.
func main() {
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

	log.Println("Building repositories...")
	userRepo := repositories.UserRepository{DB: db}

	log.Println("Building use cases...")
	authUc := brules.AuthUC{UsrRp: userRepo}

	log.Println("Building controllers...")
	authC := controllers.AuthController{AuthUC: authUc}

	log.Println("Building App...")
	app := api.App{Auth: authC}

	server := app.NewAPI()

	log.Println("Starting server on port :8000")

	log.Fatal(http.ListenAndServe(":8000", server))

}
