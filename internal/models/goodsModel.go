package models

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	StoreId  int
	Name     string `json:"name"`
	Category string `json:"category"`
	Amount   int    `json:"amount"`
}
