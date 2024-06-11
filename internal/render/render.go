package render

import (
	"bytes"
	"fmt"

	//"http/template"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/mharner33/bookings/internal/config"
	"github.com/mharner33/bookings/internal/models"
)

var app *config.AppConfig
var pathToTemplates = "./templates"
var functions = template.FuncMap{}

func NewTemplate(a *config.AppConfig) {
	app = a
}

// Returns data that should be available to all functions
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// Renders the template with name templ
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		//Create a template cache so we don't need to read from disk
		tc = app.TemplateCache
	} else {
		fmt.Println("Creating template cache.")
		tc, _ = CreateTemplateCache()
	}
	fmt.Println("Cache Lenght: ", len(tc))
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := t.Execute(buff, td)
	if err != nil {
		log.Println(err)
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the *.tmpl files from templates folder
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	fmt.Println(pages)
	//range through the .tmpl files
	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil
}
