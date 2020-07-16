package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// engine.Use(RequestInfo())

	engine.GET("/query", RequestInfo(), func(c *gin.Context) {
		fmt.Println("interrupt the middleware")
		c.JSON(404, map[string]interface{}{
			"code": 1,
			"msg":  c.FullPath(),
		})

	})
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  c.FullPath(),
		})

	})
	engine.Run(":9090")
}

func RequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		method := c.Request.Method
		fmt.Println("request path: ", path)
		fmt.Println("request method: ", method)
		fmt.Println("status code: ", c.Writer.Status())

		c.Next() //将中间件一分为二, 先执行c.Next之前的代码, 再执行内部代码, 完毕之后再执行Next之后的代码

		fmt.Println("status code: ", c.Writer.Status())
	}
}
