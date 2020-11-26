package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name         string `form:"name"`
	Password     string `form:"password"`
	PasswordHash string
	Phone        string `form:"phone" gorm:"unique;size:11"`
	Avatar       Image  `form:"avatar" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IdName       string `form:"id_name"`
	IdCard       string `form:"id_card" gorm:"unique;size:19"`
}
