package house

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func DetailView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "detail.html", nil)
}
func MyHouseView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "myhouse.html", nil)
}
func NewHouseView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "newhouse.html", nil)
}
func SearchView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "search.html", nil)
}
