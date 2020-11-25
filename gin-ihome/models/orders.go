package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID int
	HouseID int
	BeginDate time.Time
	EndDate time.Time
	Days int
	Amount float64
	Status 

	user_id = db.Column(db.Integer, db.ForeignKey("ihome_user.id"), nullable=False)
	house_id = db.Column(db.Integer, db.ForeignKey("ihome_house.id"), nullable=False)
	begin_date = db.Column(db.DateTime, nullable=False)  # 入住时间
	end_date = db.Column(db.DateTime, nullable=False)  # 离店时间
	days = db.Column(db.Integer, nullable=False)  # 住多少天
	house_price = db.Column(db.Integer, nullable=False)  # 房间价格
	amount = db.Column(db.Integer, nullable=False)  # 总价格
	status = db.Column(
	db.Enum(
	"WAIT_ACCEPT",  # 待接单,
	"WAIT_PAYMENT",  # 待支付
	"PAID",  # 已支付
	"WAIT_COMMENT",  # 待评价
	"COMPLETE",  # 已完成
	"CANCELED",  # 已取消
	"REJECTED"  # 已拒单
),
default="WAIT_ACCEPT", index=True)
	comment = db.Column(db.Text)  # 评论
}
