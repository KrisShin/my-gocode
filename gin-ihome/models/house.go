package models

import "gorm.io/gorm"

type House struct {
	gorm.Model
	// 房屋主人的用户编号
	Owner       User       `gorm:"not null;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Area        Area       `gorm:"not null;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 归属地的区域
	Title       string     `gorm:"not null"`
	Price       float64    `gorm:"not null"`
	Address     string     `gorm:"not null"`
	RoomCount   int        `gorm:"default:1"` // 房间数目
	Acreage     float64    `gorm:"default:0"` // 房屋面积
	Unit        string     // 房屋单元， 如几室几厅
	Capacity    int        `gorm:"default:1"` // 房屋容纳的人数
	Beds        int        `gorm:"default:0"` // 房屋床铺的配置
	Deposit     float64    `gorm:"default:0"` // 房屋押金
	MinDays     int        `gorm:"default:1"` // 最少入住天数
	MaxDays     int        `gorm:"default:0"` // 最多入住天数，0表示不限制
	OrdersCount int        `gorm:"default:0"` // 预订完成的该房屋的订单数
	Tenants     []User     `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Facilities  []Facility `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Images      []Image    `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Area struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type Facility struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Icon string `gorm :"not null"`
}
