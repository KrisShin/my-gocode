package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func RegisterView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func MyView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "my.html", nil)
}

func LordersView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "lorders.html", nil)
}

func ProfileView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "profile.html", nil)
}

func AuthView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "auth.html", nil)
}
