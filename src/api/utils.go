package api

import (
	"log"
	"net/http"
)

func (a App) renderHTMLPage(template string) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		// ExecuteTemplate allow us to select one of the named templates in the template set (in this case we have "index.html" and "register.html").
		// We're executing index.html because it'll render the top-level contents of this file and when it reaches the "body" block, it'll replace with the latest definition of that block in that template set.
		if err := a.pageTemplates[template].ExecuteTemplate(res, "index.html", nil); err != nil {
			log.Println(err)
			http.Error(res, "Failed to excute template", http.StatusInternalServerError)
		}
	}
}
