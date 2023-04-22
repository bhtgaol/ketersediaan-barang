package config

import "log"

// AppConfig hold the application config
type AppConfig struct {
	UseCache     bool
	InfoLog      *log.Logger
	InProduction bool
}
