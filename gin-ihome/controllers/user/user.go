package user

import (
	"gin-ihome/global"
	"gin-ihome/models"
	"gin-ihome/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	user := models.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		panic("bind data failed")
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "bind data failed"})
	}

	global.GVA_DB.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}
