package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"

	"go-breeders/pets"
)

func (app *application) ShowHome(responseWriter http.ResponseWriter, request *http.Request) {
	app.renderTemplate(responseWriter, "home.page.gohtml", nil)
}

func (app *application) ShowPage(responseWriter http.ResponseWriter, request *http.Request) {
	page := chi.URLParam(request, "page")
	app.renderTemplate(responseWriter, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(responseWriter http.ResponseWriter, request *http.Request) {
	var tool toolbox.Tools
	_ = tool.WriteJSON(responseWriter, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(responseWriter http.ResponseWriter, request *http.Request) {
	var tool toolbox.Tools
	_ = tool.WriteJSON(responseWriter, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns(responseWriter http.ResponseWriter, request *http.Request) {
	app.renderTemplate(responseWriter, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(responseWriter http.ResponseWriter, request *http.Request) {
	var tool toolbox.Tools
	dog, error := pets.NewPetFromAbstractFactory("dog")
	if error != nil {
		_ = tool.ErrorJSON(responseWriter, error, http.StatusBadRequest)
		return
	}

	_ = tool.WriteJSON(responseWriter, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(responseWriter http.ResponseWriter, request *http.Request) {
	var tool toolbox.Tools
	cat, error := pets.NewPetFromAbstractFactory("cat")
	if error != nil {
		_ = tool.ErrorJSON(responseWriter, error, http.StatusBadRequest)
		return
	}

	_ = tool.WriteJSON(responseWriter, http.StatusOK, cat)
}
