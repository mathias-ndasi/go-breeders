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

	multiplexer.Get("/", app.ShowHome)

	return multiplexer;
}