package api

import (
	"html/template"
	"net/http"
	"performance_manager/src/controllers"
)

// making each rote a method under `App` is fine for small applications.
// Must separate into one handler struct per feature, each holding only the use cases
// it needs as the app increases.
type App struct {
	pageTemplates map[string]*template.Template
	Auth          controllers.AuthController
}

func NewApp(pageTemplates map[string]*template.Template, authC controllers.AuthController) App {
	return App{pageTemplates: pageTemplates, Auth: authC}
}

func (a App) NewAPI() *http.ServeMux {
	server := http.NewServeMux()

	server.HandleFunc("GET /register", a.renderHTMLPage("register.html"))
	server.HandleFunc("POST /register", a.CreateUserHandler)

	//server.HandleFunc("/register", func(res http.ResponseWriter, req *http.Request) {
	//	// ExecuteTemplate allow us to select one of the named templates in the template set (in this case we have "index.html" and "register.html").
	//	// We're executing index.html because it'll render the top-level contents of this file and when it reaches the "body" block, it'll replace with the latest definition of that block in that template set.
	//	if err = pageTemplates["register.html"].ExecuteTemplate(res, "index.html", nil); err != nil {
	//		log.Println(err)
	//		http.Error(res, "Failed to excute template", http.StatusInternalServerError)
	//	}
	//})

	//server.HandleFunc("/account-activation", func(res http.ResponseWriter, req *http.Request) {
	//	if err = pageTemplates["account_activation.html"].ExecuteTemplate(res, "index.html", nil); err != nil {
	//		log.Println(err)
	//		http.Error(res, "Failed to excute template", http.StatusInternalServerError)
	//	}
	//})

	//server.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
	//	if err = pageTemplates["login.html"].ExecuteTemplate(res, "index.html", nil); err != nil {
	//		log.Println(err)
	//		http.Error(res, "Failed to excute template", http.StatusInternalServerError)
	//	}
	//})

	//server.HandleFunc("/reset-password", func(res http.ResponseWriter, req *http.Request) {
	//	if err = pageTemplates["reset_password.html"].ExecuteTemplate(res, "index.html", nil); err != nil {
	//		log.Println(err)
	//		http.Error(res, "Failed to execute template", http.StatusInternalServerError)
	//	}
	//})

	//server.HandleFunc("/update-password", func(res http.ResponseWriter, req *http.Request) {
	//	if err = pageTemplates["update_password.html"].ExecuteTemplate(res, "index.html", nil); err != nil {
	//		log.Println(err)
	//		http.Error(res, "Failed to execute template", http.StatusInternalServerError)
	//	}
	//})

	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return server
}
