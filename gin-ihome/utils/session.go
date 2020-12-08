package utils

import (
	. "gin-ihome/models"
	"github.com/gin-contrib/sessions"
)

func GetUser(session sessions.Session) User {
	user := session.Get("loginUser")
	if user == nil {
		return nil
	} else {
		return user.(User)
	}
}

func SetUser(session sessions.Session, user User) {
	session.Set("loginUser", user)
	session.Save()
}

func ClearLogin(session sessions.Session) {
	session.Delete("loginUser")
	session.Save()
}

func IsLogin(session sessions.Session) bool {
	user := session.Get("loginUser")
	if user == nil {
		return false
	} else {
		return true
	}
}
