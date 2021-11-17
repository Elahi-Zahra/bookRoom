package main

import (
	"github.com/Elahi-Zahra/bookRoom/pkg/config"
	"github.com/Elahi-Zahra/bookRoom/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mx := chi.NewRouter()

	mx.Use(middleware.Recoverer)
	mx.Use(NoSurf)
	mx.Use(SessionLoad)
	//mx.Use(WriteToConsole)

	fileServer := http.FileServer(http.Dir("./static/"))
	mx.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mx.Get("/", handlers.Repo.Home)
	mx.Get("/about", handlers.Repo.About)
	return mx
}
