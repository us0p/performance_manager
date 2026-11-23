package api

import (
	"net/http"
)

func (a App) CreateUserHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(res, "Error extracting form values", http.StatusBadRequest)
		return
	}

	err := a.Auth.CreateUserController(req.Context(), req.Form.Get("email"), req.Form.Get("name"), req.Form.Get("password"))
	if err != nil {
		http.Error(res, "validation error", http.StatusConflict)
		return
	}

	res.WriteHeader(http.StatusCreated)
	return
}
