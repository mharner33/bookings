package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/mharner33/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is %T", v))
	}
}
