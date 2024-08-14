package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}
