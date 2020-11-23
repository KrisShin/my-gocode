package c02

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id   int
	Name string
	Age  int
	Addr string
}

func User(con *gin.Context) {

	name := "Kris"
	con.HTML(200, "user/user.html", name)
}

// 结构体渲染
func UserInfoStruct(con *gin.Context) {
	//user_info := UserInfo{Id: 1, Name: "Shin", Age: 24, Addr: "sichuan,chengdu"}

	var user_info1 UserInfo
	user_info1.Id = 2
	user_info1.Name = "Shin2"
	user_info1.Age = 18
	user_info1.Addr = "sichuan,leshan"
	con.HTML(http.StatusOK, "cha02/user_info.html", user_info1)
}

func ArrController(con *gin.Context) {
	arr := [3]int{1, 3, 4}
	con.HTML(http.StatusOK, "cha02/arr.html", arr)
}

func ArrStruct(con *gin.Context) {
	arr_struct := [3]UserInfo{{Id: 1, Name: "kris", Age: 12, Addr: "xxx"}, {Id: 2, Name: "kris2", Age: 13, Addr: "xxx2"}, {Id: 3, Name: "kris3", Age: 14, Addr: "xxx3"}}
	con.HTML(http.StatusOK, "cha02/arr_struct.html", arr_struct)

}
