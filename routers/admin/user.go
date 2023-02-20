package admin

import (
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/app"
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	appG := app.GetGin(c)
	json := app.BindJson[request.Login](c)
	data, code := userService.AdminLogin(json.Username, json.Password, util.IP.GetIpAddress(c))
	appG.Response(code, data)
}

func GetUserList(c *gin.Context) {
	appG := app.GetGin(c)
	users, total := userService.GetUserList(app.BindValidQuery[request.PageQuery](c))
	appG.Response(e.SUCCESS, gin.H{
		"users": users,
		"total": total,
	})
	return
}
func DeleteUsers(c *gin.Context) {
	appG := app.GetGin(c)
	appG.Response(userService.DeleteUsers(app.BindJson[request.IdsJson](c)), nil)
	return
}
func UpdateUserStatus(c *gin.Context) {
	appG := app.GetGin(c)
	appG.Response(userService.UpdateUserStatus(app.BindJson[request.UpdateStatus](c)), nil)
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
