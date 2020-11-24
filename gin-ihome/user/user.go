package user

import (
	"gin-ihome/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = utils.Connection()

func LoginView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func RegisterView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func Captcha(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"captcha": utils.RandString(4)})
}

func Register(ctx *gin.Context) {
	var user User
	user.Name = ctx.Request.FormValue("name")
	user.Password = ctx.Request.FormValue("password")
	user.Phone = ctx.Request.FormValue("phone")
	//user.Avatar = ctx.Request.FormValue("avatar")
	db.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}
