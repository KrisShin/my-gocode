package c01

import "github.com/gin-gonic/gin"

func Hello(con *gin.Context) {
	name := "Kris"
	con.HTML(200, "index/index.html", name)
}

