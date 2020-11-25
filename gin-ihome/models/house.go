package models

type House struct {
	Id int `gorm:"AUTO_INCREMENT"`
	Addr string `form:"address" gorm:"not null"`
}
