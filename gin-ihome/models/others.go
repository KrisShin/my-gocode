package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Name string
	Url  string
}
