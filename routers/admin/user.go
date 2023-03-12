package admin

import (
	"github.com/WayeeeX/go-gin-example/models"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
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
func UpdateUser(c *gin.Context) {
	user := util.CopyProperties[models.User](app.BindJson[request.UpdateUserAdmin](c))
	models.Updates[models.User](&user, "id = ?", user.ID)
	util.OK(c)
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

func GetMyDetail(c *gin.Context) {
	userID, _ := c.Get("userID")
	user := models.GetOne[models.User](models.User{}, "id = ?", userID)
	util.Response(c, e.SUCCESS, gin.H{
		"user": user,
	})
}
func GetUserDetail(c *gin.Context) {
	util.Response(c, e.SUCCESS, gin.H{
		"user": userService.GetUserDetailByID(app.BindValidQuery[request.IdQuery](c).Id),
	})
	return
}
func GetUserSelectList(c *gin.Context) {
	req := app.BindValidQuery[request.PageQuery](c)
	var total uint64
	var users []response.UserSelect
	err := models.DB.Table("tb_user").Select("id,nickname,avatar").Where("nickname like ?", "%"+req.Keyword+"%").Count(&total).Limit(req.PageSize).Offset(util.GetOffset(req)).Scan(&users).Error
	if err != nil {
		panic(err)
	}
	util.Response(c, e.SUCCESS, gin.H{
		"users": users,
		"total": total,
	})
	return
}
