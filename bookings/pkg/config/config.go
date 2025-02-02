package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// Appconfig holds applocation configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
