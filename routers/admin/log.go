package admin

import (
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/app"
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetLoginLogList(c *gin.Context) {
	req := app.BindValidQuery[request.PageQuery](c)
	util.Response(c, e.SUCCESS, logLoginService.GetList(req))
	return
}
