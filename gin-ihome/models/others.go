package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Url       string
	HouseID   int
	CommentID int
}
