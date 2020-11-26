package house

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
