package dao

import (
	"gin-ihome/config"
	"gin-ihome/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := config.DB_USER + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME + "?" + config.DB_CONFIG
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func MigrateDB(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.Facility{},
		&models.Area{},
		&models.Image{},
		&models.User{},
		&models.House{},
		&models.Order{},
		)
}

func CloseConn(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

func Initialize(db *gorm.DB) {
	facilities := []models.Facility{
		models.Facility{Name: "空调"},
		models.Facility{Name: "电视"},
		models.Facility{Name: "冰箱"},
		models.Facility{Name: "电风扇"},
		models.Facility{Name: "无线网"},
		models.Facility{Name: "宽带"},
		models.Facility{Name: "燃气灶"},
		models.Facility{Name: "电磁炉"},
		models.Facility{Name: "浴缸"},
		models.Facility{Name: "马桶"},
		models.Facility{Name: "锅"},
		models.Facility{Name: "电饭煲"},
		models.Facility{Name: "衣柜"},
		models.Facility{Name: "书桌"},
		models.Facility{Name: "椅子"},
		models.Facility{Name: "茶几"},
		models.Facility{Name: "沙发"},
		models.Facility{Name: "洗衣机"},
	}
	db.Create(&facilities)

	areas := []models.Area{
		models.Area{Name: "东城区"},
		models.Area{Name: "西城区"},
		models.Area{Name: "朝阳区"},
		models.Area{Name: "海淀区"},
		models.Area{Name: "昌平区"},
		models.Area{Name: "丰台区"},
		models.Area{Name: "房山区"},
		models.Area{Name: "通州区"},
		models.Area{Name: "顺义区"},
		models.Area{Name: "大兴区"},
		models.Area{Name: "怀柔区"},
		models.Area{Name: "平谷区"},
		models.Area{Name: "密云区"},
		models.Area{Name: "延庆区"},
		models.Area{Name: "石景山区"},
		models.Area{Name: "门头沟区"},
	}
	db.Create(&areas)
}
