package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookingView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "booking.html", nil)
}

func OrdersView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "orders.html", nil)
}
