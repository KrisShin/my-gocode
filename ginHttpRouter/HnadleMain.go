package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// http://localhost:9090/hello?name=kris
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		name := context.DefaultQuery("name", "hello")
		fmt.Println(name)

		// 输出
		context.Writer.Write([]byte("Hello ," + name))
	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		username, _ := context.GetPostForm("username")
		fmt.Println(username)

		password, _ := context.GetPostForm("password")
		fmt.Println(password)

		context.Writer.Write([]byte("login user:" + username + "pwd:" + password))
	})

	engine.Run(":9090")
}
