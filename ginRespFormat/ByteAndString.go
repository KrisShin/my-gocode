package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){
	engine := gin.Default()

	engine.GET("/helloByte", func (context *gin.Context)  {
		fullPath := "请求路径:"+ context.FullPath()
		fmt.Println(fullPath)
		context.Writer.Write([]byte(fullPath))
	})

	engine.GET("/helloString", func (context *gin.Context)  {
		fullPath := "请求路径:"+ context.FullPath()
		fmt.Println(fullPath)
		context.Writer.WriteString(fullPath)
	})
	engine.Run()
}