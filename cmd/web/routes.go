package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mharner33/bookings/pkg/config"
	"github.com/mharner33/bookings/pkg/handlers"
)

// Using Chi router as http lib doesn't support middleware
func routes(app *config.AppConfig) http.Handler {
	// mux := http.NewServeMux()
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	// mux.HandleFunc("GET /", http.HandlerFunc(handlers.Repo.Home))
	// mux.HandleFunc("GET /about", http.HandlerFunc(handlers.Repo.About))

	// FileServer serves static files from the static directory
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
