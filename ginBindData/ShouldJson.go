package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.POST("/addstudent", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		var per Person
		if err := c.BindJSON(&per); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(per.Name)
		fmt.Println(per.Gender)
		fmt.Println(per.Age)

		c.Writer.Write([]byte(per.Name + per.Gender + "register"))

	})
	engine.Run()
}

type Person struct {
	Name   string `form:"name"`
	Gender string `form:"gender"`
	Age    int    `form:"age"`
}
