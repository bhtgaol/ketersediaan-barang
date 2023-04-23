package config

import (
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig hold the application config
type AppConfig struct {
	UseCache     bool
	InfoLog      *log.Logger
	InProduction bool
	Session      *scs.SessionManager
}

// Repository is repository type
type Repository struct {
	App *AppConfig
}
