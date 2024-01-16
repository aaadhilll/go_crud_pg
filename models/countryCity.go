package models

import "gorm.io/gorm"

type CountryCity struct {
	gorm.Model
	Country string
	City  string
}