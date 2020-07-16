package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	routerGroup := engine.Group("/user")

	routerGroup.POST("/register", registerHandle)

	routerGroup.POST("/login", loginHandle)

	routerGroup.DELETE("/:id", deleteHandle)

	engine.Run()
}

func registerHandle(c *gin.Context) {
	fullPath := "user register " + c.FullPath()
	fmt.Println(fullPath)

	c.Writer.WriteString(fullPath)
}

func loginHandle(c *gin.Context) {
	fullPath := "user login " + c.FullPath()
	fmt.Println(fullPath)

	c.Writer.WriteString(fullPath)
}

func deleteHandle(c *gin.Context) {
	fullPath := "user delete " + c.FullPath()
	userID := c.Param("id")
	fmt.Println(fullPath + " " + userID)

	c.Writer.WriteString(fullPath)
}
