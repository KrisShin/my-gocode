package main

import (
	"gin_shop_project/c01"
	"gin_shop_project/c02"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/**/*")
	router.Static("/static", "static")
	//router.LoadHTMLFiles("index.html", "news.html")
	router.GET("/", c01.Hello)
	router.GET("/user", c02.User)
	router.GET("/user_struct", c02.UserInfoStruct)
	router.GET("/arr", c02.ArrController)
	router.GET("/arr_struct", c02.ArrStruct)
	router.GET("/map",c02.MapController)
	router.GET("/map_struct",c02.MapStruct)

	// 路由获取参数 强制匹配
	router.GET("/param1/:id",c02.Param1)
	// 路由匹配参数 非强制匹配
	router.GET("/param2/*id",c02.Param2)

	router.GET("/query",c02.GetQueryData)
	router.GET("/to_user_add",c02.ToUserAdd)
	router.GET("/do_user_add",c02.DoUserAdd)

	router.Run(":8989")
}
