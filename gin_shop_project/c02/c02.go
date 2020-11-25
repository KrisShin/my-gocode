package c02

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Id   int
	Name string `form:"username"`
	Age  int    `form:"age"`
	Addr string `form:"addr"`
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

func MapController(ctx *gin.Context) {
	map_data := map[string]string{"name": "kris", "age": "18"}
	map_data2 := map[string]int{"age": 19}
	map_data3 := map[string]interface{}{"map_data": map_data, "map_data2": map_data2}
	ctx.HTML(http.StatusOK, "cha02/map.html", map_data3)
}

func MapStruct(ctx *gin.Context) {
	map_struct := map[string]UserInfo{"user": {Id: 1, Name: "kris", Age: 19, Addr: "xxxk"}}
	ctx.HTML(http.StatusOK, "cha02/map_struct", map_struct)
}

func Param1(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello %s", id)
}
func Param2(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello %s", id)
}
func GetQueryData(ctx *gin.Context) {
	id := ctx.Query("id")

	name := ctx.DefaultQuery("name", "www")
	fmt.Println(name)
	ctx.String(http.StatusOK, "hello %s", id)
}

func ToUserAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "cha02/user_add", nil)
}

func DoUserAdd(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username, password)
	ctx.String(http.StatusOK, "success!!!")
}
