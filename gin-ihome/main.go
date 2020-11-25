package main

import (
	"gin-ihome/controllers/house"
	"gin-ihome/controllers/user"
	"gin-ihome/dao"
	_ "gin-ihome/dao"
	"gin-ihome/global"
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
	global.GVA_DB = dao.Connection()
	dao.MigrateDB(global.GVA_DB)
	//dao.Initialize(global.GVA_DB)
	defer dao.CloseConn(global.GVA_DB)

	//engine := gin.Default()
	//engine.Static("/static", "static")
	//engine.LoadHTMLGlob("templates/*")
	//
	//MainView(engine)
	//UsersUrl(engine)
	//HouseUrl(engine)
	//OrdersUrl(engine)
	//
	//engine.Run(":9000")
}
