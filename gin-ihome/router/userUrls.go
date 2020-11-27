package router

import (
	"gin-ihome/controllers/user"
	"github.com/gin-gonic/gin"
)

// 用户模块接口
func UsersUrl(engine *gin.Engine) {
	users := engine.Group("/users")
	users.POST("/captcha", user.Captcha)
	users.POST("/register", user.Register)
	users.POST("/login", user.Login)
}