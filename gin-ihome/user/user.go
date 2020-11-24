package user

import (
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
	captcha := map[string]string{"captcha": utils.RandString(4)}
	ctx.JSON(http.StatusOK, captcha)
}

func Register(ctx *gin.Context) {

}
