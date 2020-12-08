package user

import (
	"fmt"
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

	if err := global.DB.Find(&user, "phone = ?", user.Phone).Error; err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": 1003, "msg": "该手机号已注册,请登陆"})
		return
	}

	user.PasswordHash = utils.GetMD5(user.Password)

	global.DB.Select(`created_at`, `updated_at`, `deleted_at`, `name`, `password`, `password_hash`, `phone`).Create(&user)
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
	fmt.Println(phone, password, passwordHash)
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
	ok := utils.IsLogin(session)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg":"请重新登陆"})
		return
	}
	utils.ClearLogin(session)
	ctx.JSON(http.StatusOK, gin.H{"code": 200})
}

func Auth(ctx *gin.Context) {
	session := sessions.Default(ctx)
	ok := utils.IsLogin(session)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg":"请重新登陆"})
		return
	}
	user := utils.GetUser(session)
	idName := ctx.PostForm("id_name")
	idCard := ctx.PostForm("id_card")
	idRex := regexp.MustCompile(`^\d{17}[\d,x,X]$`)

	if id := idRex.FindAllString(idCard, -1); id == nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 1005, "msg": "身份证有误"})
		return
	}

	global.DB.Model(&user).UpdateColumns(User{IdName: idName, IdCard: idCard})
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func UpdateUser(ctx *gin.Context) {
	category := ctx.Param("category")
	if category == "pwd" {

	} else if category == "info" {
		var user User


		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusAccepted, gin.H{"code": 1001, "msg": "请完整填写数据"})
			return
		}
	}
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

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
//func Login(c *gin.Context) {
//	var L request.Login
//	_ = c.ShouldBindJSON(&L)
//	if err := utils.Verify(L, utils.LoginVerify); err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	if store.Verify(L.CaptchaId, L.Captcha, true) {
//		U := &model.SysUser{Username: L.Username, Password: L.Password}
//		if err, user := service.Login(U); err != nil {
//			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
//			response.FailWithMessage("用户名不存在或者密码错误", c)
//		} else {
//			tokenNext(c, *user)
//		}
//	} else {
//		response.FailWithMessage("验证码错误", c)
//	}
//}

// 登录以后签发jwt
//func tokenNext(c *gin.Context, user model.SysUser) {
//	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
//	claims := request.CustomClaims{
//		UUID:        user.UUID,
//		ID:          user.ID,
//		NickName:    user.NickName,
//		Username:    user.Username,
//		AuthorityId: user.AuthorityId,
//		BufferTime:  60 * 60 * 24, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
//		StandardClaims: jwt.StandardClaims{
//			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
//			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 7天
//			Issuer:    "qmPlus",                       // 签名的发行者
//		},
//	}
//	token, err := j.CreateToken(claims)
//	if err != nil {
//		global.GVA_LOG.Error("获取token失败", zap.Any("err", err))
//		response.FailWithMessage("获取token失败", c)
//		return
//	}
//	if !global.GVA_CONFIG.System.UseMultipoint {
//		response.OkWithDetailed(response.LoginResponse{
//			User:      user,
//			Token:     token,
//			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
//		}, "登录成功", c)
//		return
//	}
//	if err, jwtStr := service.GetRedisJWT(user.Username); err == redis.Nil {
//		if err := service.SetRedisJWT(token, user.Username); err != nil {
//			global.GVA_LOG.Error("设置登录状态失败", zap.Any("err", err))
//			response.FailWithMessage("设置登录状态失败", c)
//			return
//		}
//		response.OkWithDetailed(response.LoginResponse{
//			User:      user,
//			Token:     token,
//			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
//		}, "登录成功", c)
//	} else if err != nil {
//		global.GVA_LOG.Error("设置登录状态失败", zap.Any("err", err))
//		response.FailWithMessage("设置登录状态失败", c)
//	} else {
//		var blackJWT model.JwtBlacklist
//		blackJWT.Jwt = jwtStr
//		if err := service.JsonInBlacklist(blackJWT); err != nil {
//			response.FailWithMessage("jwt作废失败", c)
//			return
//		}
//		if err := service.SetRedisJWT(token, user.Username); err != nil {
//			response.FailWithMessage("设置登录状态失败", c)
//			return
//		}
//		response.OkWithDetailed(response.LoginResponse{
//			User:      user,
//			Token:     token,
//			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
//		}, "登录成功", c)
//	}
//}

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
//func Register(c *gin.Context) {
//	var R request.Register
//	_ = c.ShouldBindJSON(&R)
//	if err := utils.Verify(R, utils.RegisterVerify); err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	user := &model.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
//	err, userReturn := service.Register(*user)
//	if err != nil {
//		global.GVA_LOG.Error("注册失败", zap.Any("err", err))
//		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
//	} else {
//		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
//	}
//}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
//func ChangePassword(c *gin.Context) {
//	var user request.ChangePasswordStruct
//	_ = c.ShouldBindJSON(&user)
//	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	U := &model.SysUser{Username: user.Username, Password: user.Password}
//	if err, _ := service.ChangePassword(U, user.NewPassword); err != nil {
//		global.GVA_LOG.Error("修改失败", zap.Any("err", err))
//		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
//	} else {
//		response.OkWithMessage("修改成功", c)
//	}
//}
