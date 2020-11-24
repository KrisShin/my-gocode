package house

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
