package user

import (
	"errors"
	. "gin-ihome/config"
	"gin-ihome/global"
	. "gin-ihome/models"
	"gin-ihome/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func Captcha(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"captcha": utils.RandString(4)})
}

func Register(ctx *gin.Context) {
	user := User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1001, "msg": "请完整填写数据"})
		return
	}

	phoneRex := regexp.MustCompile(`^1[345789]\d{9}$`)
	if phone := phoneRex.FindAllString(user.Phone, -1); phone == nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1002, "msg": "手机号码格式错误"})
		return
	}

	if !errors.Is(global.DB.First(&user, "phone = ?", user.Phone).Error, ErrNotFound) {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1003, "msg": "该手机号已注册,请登陆"})
		return
	}

	user.PasswordHash = utils.GetMD5(user.Password)

	err := global.DB.Select(`created_at`, `updated_at`, `deleted_at`, `name`, `password`, `password_hash`, `phone`).Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1007, "msg": "数据库错误", "error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func Login(ctx *gin.Context) {
	var user User
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	phoneRex := regexp.MustCompile(`^1[345789]\d{9}$`)
	if phone := phoneRex.FindAllString(phone, -1); phone == nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1002, "msg": "手机号码格式错误"})
		return
	}

	if err := global.DB.First(&user, "phone = ?", phone).Error; err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1003, "msg": "该用户不存在,请注册"})
		return
	}
	passwordHash := utils.GetMD5(password)
	if err := global.DB.Where("Phone = ? AND Password_Hash = ?", phone, passwordHash).First(&user).Error; err == nil {
		session := sessions.Default(ctx)
		if session.Get("loginUser") == nil {
			utils.SetUser(session, user)
		}
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
		return
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1004, "msg": "密码错误"})
		return
	}
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	_, ok := utils.GetUser(session)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "没有登陆"})
		return
	}
	utils.ClearLogin(session)
	ctx.JSON(http.StatusOK, gin.H{"code": 200})
}

func Update(ctx *gin.Context) {
	category := ctx.DefaultQuery("cate", "err")
	session := sessions.Default(ctx)
	user, ok := utils.GetUser(session)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "请重新登陆"})
		return
	}
	if category == "pwd" {
		passwordHash := utils.GetMD5(ctx.PostForm("password"))
		if user.PasswordHash != passwordHash {
			ctx.JSON(http.StatusAccepted, gin.H{"code": 1004, "msg": "密码错误"})
			return
		}

		newPwd := ctx.PostForm("newPassword")
		newPwdHash := utils.GetMD5(newPwd)
		if passwordHash == newPwdHash {
			ctx.JSON(http.StatusAccepted, gin.H{"code": 1006, "msg": "新密码不能与原密码相同"})
			return
		}

		global.DB.Model(&user).UpdateColumns(User{Password: newPwd, PasswordHash: newPwdHash})
		utils.ClearLogin(session)
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	} else if category == "info" {

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusAccepted, gin.H{"code": 1001, "msg": "请完整填写数据"})
			return
		}
	} else if category == "auth" {
		idName := ctx.PostForm("id_name")
		idCard := ctx.PostForm("id_card")
		idRex := regexp.MustCompile(`^\d{17}[\d,x,X]$`)

		if user.IdCard != "" || user.IdName != "" {
			ctx.JSON(http.StatusOK, gin.H{"code": 1006, "msg": "实名认证信息不可修改"})
			return
		}
		if id := idRex.FindAllString(idCard, -1); id == nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 1005, "msg": "身份证有误"})
			return
		}

		user.IdName = idName
		user.IdCard = idCard
		global.DB.Model(&user).UpdateColumns(User{IdName: idName, IdCard: idCard})
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1001, "msg": "请完整填写数据"})
		return
	}
	utils.UpdateUser(session, user.ID)
}

func DeleteUser(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if user := global.DB.First(&User{}, "ID = ?", userId); user != nil {
		global.DB.Delete(&user)
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 1003, "msg": "用户不存在"})
	}
}
