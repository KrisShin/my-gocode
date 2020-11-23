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

	router.Run(":8989")
}
