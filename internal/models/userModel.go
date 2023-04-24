package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	StatusId     int
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}
