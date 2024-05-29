package render

import (
	"bytes"
	"fmt"

	//"http/template"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/mharner33/bookings/pkg/config"
	"github.com/mharner33/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

// Returns data that should be available to all functions
func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// Renders the template with name templ
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
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

	td = AddDefaultData(td)

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
	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	fmt.Println(pages)
	//range through the .tmpl files
	for _, page := range pages {
		name := filepath.Base(page)
		//fmt.Println(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil
}
