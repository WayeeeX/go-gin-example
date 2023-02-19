package admin

import (
	"github.com/EDDYCJY/go-gin-example/models/request"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	appG := app.GetGin(c)
	var json request.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	data, code := userService.AdminLogin(json.Username, json.Password, util.IP.GetIpAddress(c))
	appG.Response(code, data)
}

func GetUserList(c *gin.Context) {
	appG := app.GetGin(c)
	appG.Response(userService.GetUserList(c), nil)
	return
}
func CheckExistUsername(c *gin.Context) {
	appG := app.GetGin(c)
	username, ok := c.GetQuery("username")
	if !ok {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	exist := userService.CheckUserExistByUsername(username)
	if exist {
		appG.Response(e.ERROR_USERNAME_EXIST, nil)
		return
	}
	appG.Response(e.SUCCESS, nil)
}

func CheckExistNickname(c *gin.Context) {
	appG := app.GetGin(c)
	nickname, ok := c.GetQuery("nickname")
	if !ok {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	exist := userService.CheckUserExistByNickname(nickname)
	if exist {
		appG.Response(e.ERROR_NICKNAME_EXIST, nil)
		return
	}
	appG.Response(e.SUCCESS, nil)
}

func GetUserDetail(c *gin.Context) {
	appG := app.GetGin(c)
	userID, _ := c.Get("userID")
	user := userService.GetUserDetailByID(userID.(uint))
	appG.Response(e.SUCCESS, user)
}
