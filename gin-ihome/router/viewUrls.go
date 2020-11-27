package router

import (
	"gin-ihome/controllers/house"
	"gin-ihome/controllers/user"
	"github.com/gin-gonic/gin"
)

// 所有页面路由
func MainView(engine *gin.Engine) {
	engine.GET("/", house.IndexView)
	engine.GET("/login", user.LoginView)
	engine.GET("/register", user.RegisterView)
}
