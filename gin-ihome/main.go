package main

import (
	"gin-ihome/dao"
	"gin-ihome/global"
	. "gin-ihome/router"
	"github.com/gin-gonic/gin"
)

func main() {
	global.GVA_DB = dao.Connection()
	dao.MigrateDB(global.GVA_DB)
	//dao.Initialize(global.GVA_DB)
	db, _ := global.GVA_DB.DB()
	defer db.Close()

	engine := gin.Default()
	engine.Static("/static", "static")
	engine.LoadHTMLGlob("templates/*")

	MainView(engine)
	UsersUrl(engine)
	HouseUrl(engine)
	OrdersUrl(engine)

	engine.Run(":9000")
}
