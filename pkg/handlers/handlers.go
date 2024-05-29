package handlers

import (
	"net/http"

	"github.com/mharner33/bookings/pkg/config"
	"github.com/mharner33/bookings/pkg/models"
	"github.com/mharner33/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Create a new repository to easily swap out config options
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) CityView(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "city-view-suite.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Honeymoon(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "honeymoon-suite.page.tmpl", &models.TemplateData{})
}
