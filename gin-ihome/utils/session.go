package utils

import (
	"gin-ihome/global"
	. "gin-ihome/models"
	"github.com/gin-contrib/sessions"
)

func GetUser(session sessions.Session) (User, bool) {
	user, ok := session.Get("loginUser").(User)
	return user, ok
}

func SetUser(session sessions.Session, user User) {
	session.Set("loginUser", user)
	session.Save()
}

func UpdateUser(session sessions.Session, userId uint) {
	var user User
	global.DB.First(&user, userId)
	session.Set("loginUser", user)
}

func ClearLogin(session sessions.Session) {
	session.Delete("loginUser")
	session.Save()
}
