package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// http://localhost:9090/hello?name=kris&classes=network
	engine.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		var stu Student
		err := c.ShouldBindQuery(&stu)

		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(stu.Name)
		fmt.Println(stu.Classes)
		c.Writer.Write([]byte("hello, " + stu.Name))
	})
	engine.Run(":9090")
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}
