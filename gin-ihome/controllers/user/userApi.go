package user

import (
	"gin-ihome/global"
	"gin-ihome/models"
	"gin-ihome/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = global.GVA_DB

func Captcha(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"captcha": utils.RandString(4)})
}

func Register(ctx *gin.Context) {
	user := models.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		panic("bind data failed")
		ctx.JSON(1000, gin.H{"msg": "请完整填写数据"})
		return
	}

	err = db.Find(&user, "phone = ?", user.Phone).Error
	if err == nil {
		ctx.JSON(1001, gin.H{"msg": "用户已存在,请登陆"})
		return
	}

	user.PasswordHash = utils.GetMD5(user.Password)

	db.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}
