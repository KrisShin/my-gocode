package router

import (
	"gin-ihome/controllers/house"
	"gin-ihome/controllers/order"
	"gin-ihome/controllers/user"
	"github.com/gin-gonic/gin"
)

// 所有页面路由
func MainView(engine *gin.Engine) {
	engine.GET("/", house.IndexView)
	engine.GET("/login", user.LoginView)
	engine.GET("/register", user.RegisterView)
	engine.GET("/my", user.MyView)
	engine.GET("/myhouse", house.MyHouseView)
	engine.GET("/newhouse", house.NewHouseView)
	engine.GET("/detail", house.DetailView)
	engine.GET("/search", house.SearchView)
	engine.GET("/lorders", user.LordersView)
	engine.GET("/auth", user.AuthView)
	engine.GET("/profile", user.ProfileView)
	engine.GET("/booking", order.BookingView)
	engine.GET("/orders", order.OrdersView)
}
