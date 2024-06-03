package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mharner33/bookings/internal/config"
	"github.com/mharner33/bookings/internal/handlers"
)

// Using Chi router as http lib doesn't support middleware
func routes(app *config.AppConfig) http.Handler {
	// mux := http.NewServeMux()
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/honeymoon-suite", handlers.Repo.Honeymoon)
	mux.Get("/city-view-suite", handlers.Repo.CityView)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/make-reservation", handlers.Repo.Reservation)

	// FileServer serves static files from the static directory
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
