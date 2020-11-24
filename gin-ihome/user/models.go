package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id           int    `db:"ID" json:"id" gorm:"AUTO_INCREMENT"`
	Name         string `db:"name" json:"name"`
	Password     string `db:"password" json:"password"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
	Phone        string `db:"phone" json:"phone" gorm:"unique"`
	Avatar       string `db:"avatar" json:"avatar"`
	IdName       string `db:"id_name" json:"id_name"`
	IdCard       string `db:"id_card" json:"id_card"`

	//HouseId int `gorm:"index"`
	//orders		order.Order
}
