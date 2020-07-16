package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// 设置html目录
	engine.LoadHTMLGlob("./html/*")

	// 设置静态资源目录
	engine.Static("/img", "./img")
	engine.GET("/hellohtml", func(context *gin.Context) {
		fullPath := "请求路径:" + context.FullPath()
		fmt.Println(fullPath)

		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullPath": fullPath,
			"title":    "hello kris",
		})
	})
	engine.Run()
}
