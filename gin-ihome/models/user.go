package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `form:"name"`
	Password     string `form:"password"`
	PasswordHash string
	Phone        string `form:"phone" gorm:"unique;size:11"`
	AvatarID     int
	Avatar       Image  `form:"avatar" gorm:"foreignKey:AvatarID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IdName       string `form:"id_name"`
	IdCard       string `form:"id_card" gorm:"unique;size:19"`

	Houses []*House `gorm:"many2many:user_houses"`
	Orders []Order
}
