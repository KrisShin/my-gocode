package user

import (
	"fmt"
	"gin-ihome/dao"
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
	//user := models.User{Name: "zhoucun", Phone: "12312312344", Password: "z123123"}
	//err := ctx.ShouldBind(&user)
	//if err != nil {
	//	panic("bind data failed")
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "bind data failed"})
	//}

	fmt.Println(dao.DB)
	//ctx.JSON(http.StatusOK, gin.H{"success": true, "user": user})
	ctx.String(http.StatusOK, "success")
}
