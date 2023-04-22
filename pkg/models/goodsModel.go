package models

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name     string
	Category string
	Amount   int
}
