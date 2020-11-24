package main

import (
	"fmt"
	"gin-ihome/house"
	"gin-ihome/user"
	"gin-ihome/utils"
	"github.com/gin-gonic/gin"
)

// 所有页面路由
func MainView(engine *gin.Engine) {
	engine.GET("/", house.IndexView)
	engine.GET("/login", user.LoginView)
	engine.GET("/register", user.RegisterView)
}

// 用户模块接口
func UsersUrl(engine *gin.Engine) {
	users := engine.Group("/users")
	users.POST("/captcha", user.Captcha)
	users.POST("/register", user.Register)
}
func HouseUrl(engine *gin.Engine) {
	//house := engine.Group("/house")
	//house.POST("/captcha", user.GetCaptcha)
}
func OrdersUrl(engine *gin.Engine) {
	//orders := engine.Group("/orders")
	//users.POST("/captcha", user.GetCaptcha)
}

func main() {
	db1 := utils.Connection()
	db, _ := db1.DB()
	defer db.Close()
	if (db.HasTable(&user.User{})) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&user.User{})
		fmt.Println("created user table")
	}
	if (db.HasTable(&house.House{})) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&house.House{})
		fmt.Println("created house table")
	}
	engine := gin.Default()
	engine.Static("/static", "static")
	engine.LoadHTMLGlob("templates/*")
	MainView(engine)
	UsersUrl(engine)
	HouseUrl(engine)
	OrdersUrl(engine)
	engine.Run(":9000")
}
