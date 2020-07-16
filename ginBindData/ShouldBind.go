package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func  main()  {
	engine:=gin.Default()
	engine.POST("/register", func(c *gin.Context){
		fmt.Println(c.FullPath())

		var r Register
		if err:=c.ShouldBind(&r); err!=nil{
			log.Fatal(err.Error())
			return
		}

		fmt.Println(r.UserName)
		fmt.Println(r.Password)
		fmt.Println(r.Phone)

		c.Writer.Write([]byte(r.UserName+r.Phone+r.Password+"---Register"))
	})
	engine.Run(":9090")
}

type Register struct{
	UserName string `form:"name"`
	Phone	 string `form:"phone"`
	Password string `form:"password"`
}