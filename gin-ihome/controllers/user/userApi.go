package user

import (
	"gin-ihome/global"
	"gin-ihome/models"
	"gin-ihome/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func Captcha(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"captcha": utils.RandString(4)})
}

func Register(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1001, "msg": "请完整填写数据"})
		return
	}

	phoneRex := regexp.MustCompile(`^1[345789]\d{9}$`)
	if phone := phoneRex.FindAllString(user.Phone, -1); phone == nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1002, "msg": "手机号码格式错误"})
		return
	}

	if err := global.GVA_DB.Find(&user, "phone = ?", user.Phone).Error; err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1003, "msg": "该手机号已注册,请登陆"})
		return
	}

	user.PasswordHash = utils.GetMD5(user.Password)

	global.GVA_DB.Save(&user)
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}
