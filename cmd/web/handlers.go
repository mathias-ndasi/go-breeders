package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) ShowHome(responseWriter http.ResponseWriter, request *http.Request){
	app.renderTemplate(responseWriter, "home.page.gohtml", nil)
}

func (app *application) ShowPage(responseWriter http.ResponseWriter, request *http.Request){
	page := chi.URLParam(request, "page")
	app.renderTemplate(responseWriter, fmt.Sprintf("%s.page.gohtml", page), nil)
}
