package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		name := c.Query("name")
		fmt.Println(name)

		c.Writer.Write([]byte("hello ," + name))
	})

	engine.POST("/login", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		// GetPostForm 只能用于form表单提交的数据
		username, exist := c.GetPostForm("username")
		if exist {
			fmt.Println(username)
		}

		password, exsit := c.GetPostForm("password")
		if exsit {
			fmt.Println(password)
		}

		c.Writer.Write([]byte("Hello ," + username + "," + password))
	})

	engine.DELETE("/user/:id", func(c *gin.Context) {
		userID := c.Param("id")
		fmt.Println(userID)
		c.Writer.Write([]byte("delete 用户:" + userID))

	})

	engine.Run(":9090")
}
