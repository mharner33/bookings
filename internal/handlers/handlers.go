package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mharner33/bookings/internal/config"
	"github.com/mharner33/bookings/internal/models"
	"github.com/mharner33/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s, end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) CityView(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "city-view-suite.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Honeymoon(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "honeymoon-suite.page.tmpl", &models.TemplateData{})
}
