package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	multiplexer := chi.NewRouter();

	multiplexer.Use(middleware.Recoverer);
	multiplexer.Use(middleware.Timeout(60 * time.Second));

	fileServer := http.FileServer(http.Dir("./static"))
	multiplexer.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	
	multiplexer.Get("/", app.ShowHome)
	multiplexer.Get("/{page}", app.ShowPage)

	return multiplexer;
}