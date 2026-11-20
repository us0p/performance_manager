package api

import (
	"net/http"
	"performance_manager/src/controllers"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

type App struct {
	Auth controllers.AuthController
}

func (a App) NewAPI() *http.ServeMux {
	server := http.NewServeMux()

	api := humago.New(server, huma.DefaultConfig("Performance Manager", "0.1.0"))

	huma.Register(api, createUserOp, a.CreateUserHandler)

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
