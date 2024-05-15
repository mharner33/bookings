package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// Middleware to help prevent CSRF attacks
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Saves and Loads the session on each request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
