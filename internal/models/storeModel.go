package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name    string
	Street  string
	City    string
	State   string
	Country string
}
