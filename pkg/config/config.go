package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Infolog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
