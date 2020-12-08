package utils

import (
	. "gin-ihome/models"
	"github.com/gin-contrib/sessions"
)

func GetUser(session sessions.Session) interface{} {
	user := session.Get("loginUser")
	if user != nil {
		return user.(User)
	}
	return nil
}

func SetUser(session sessions.Session, user User) {
	session.Set("loginUser", user)
	session.Save()
}

func ClearLogin(session sessions.Session) {
	session.Delete("loginUser")
	session.Save()
}
