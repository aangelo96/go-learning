package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int
	Status      int
	WebId       uint
	Image       string
	Order       int
	Group       string
}
