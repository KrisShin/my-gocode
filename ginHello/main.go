package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径: ", context.FullPath())
		context.Writer.Write([]byte("hello gin\n"))
	})

	if err := engine.Run("0.0.0.0:9090"); err != nil {
		log.Fatal(err.Error())
	}
}
