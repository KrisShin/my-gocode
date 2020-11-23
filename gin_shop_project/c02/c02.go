package c02

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Id int

}

func User(con *gin.Context) {

	name := "Kris"
	con.HTML(200, "user/user.html", name)
}
