package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		name := context.DefaultQuery("name", "hello")
		fmt.Println(name)

		// 输出
		context.Writer.Write([]byte("Hello ," + name))
	})

	//post
	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		// username := context.PostForm("username")
		// password := context.PostForm("password")
		fmt.Println(context.PostForm("u"))
		fmt.Println(context.PostForm("p"))

		// context.Writer.Write([]byte(username + "logined"))
	})

	engine.Run(":9999")
}
