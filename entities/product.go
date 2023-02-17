package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       uint   `json:"price"`
}
