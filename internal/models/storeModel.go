package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserId  int
	Name    string `json:"name"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"counttry"`
}
