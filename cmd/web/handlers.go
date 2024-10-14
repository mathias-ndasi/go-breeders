package main

import (
	"net/http"
)

func (app *application) ShowHome(responseWriter http.ResponseWriter, request *http.Request){
	app.renderTemplate(responseWriter, "home.page.gohtml", nil)
}