package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	User      User
	House     House
	BeginDate time.Time
	EndDate   time.Time
	Days      int
	Amount    float64
	Status    string `gorm:"check:Status in ('WAIT_ACCEPT','WAIT_PAYMENT','PAID','WAIT_COMMENT','COMPLETE','CANCELED','REJECTED');default:'WAIT_ACCEPT';index"`

	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content   string
	LikeCount int
}
