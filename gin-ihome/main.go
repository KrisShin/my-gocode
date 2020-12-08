package main

import (
	"encoding/gob"
	"gin-ihome/dao"
	"gin-ihome/global"
	"gin-ihome/models"
	. "gin-ihome/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	global.DB = dao.Connection()
	dao.MigrateDB(global.DB)
	//dao.Initialize(global.DB)
	db, _ := global.DB.DB()
	defer db.Close()

	gob.Register(models.User{})
	engine := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("ihomeSession", store))
	engine.Static("/static", "static")
	engine.LoadHTMLGlob("templates/*")

	MainView(engine)
	UsersUrl(engine)
	HouseUrl(engine)
	OrdersUrl(engine)

	engine.Run(":9000")
}
