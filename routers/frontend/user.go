package frontend

import (
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/app"
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	appG := app.GetGin(c)
	var json request.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	data, code := userService.Login(json.Username, json.Password, util.IP.GetIpAddress(c))
	appG.Response(code, data)
}

func Register(c *gin.Context) {
	appG := app.GetGin(c)
	var json request.Register
	if err := c.ShouldBindJSON(&json); err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	_, code := userService.Register(json.Username, json.Password, json.Nickname)
	appG.Response(code, nil)
	return
}
