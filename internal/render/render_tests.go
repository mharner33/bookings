package render

import (
	"net/http"
	"testing"

	"github.com/mharner33/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	// AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("Flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.UseCache = true
	app.TemplateCache = tc

	var w myWriter
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	path := "home.page.tmpl"
	err = RenderTemplate(&w, r, path, &models.TemplateData{})
	if err != nil {
		t.Error("error should be nil")
	}

	path = "non-existent.page.tmpl"
	err = RenderTemplate(&w, r, path, &models.TemplateData{})
	if err == nil {
		t.Error("template should not exist")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/test-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestNewTemplate(t *testing.T) {
	NewTemplate(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
