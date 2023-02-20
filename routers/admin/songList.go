package admin

import (
	"github.com/WayeeeX/go-gin-example/models"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/app"
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

func CreateSongList(c *gin.Context) {
	appG := app.GetGin(c)
	var json request.CreateSongList
	if err := c.ShouldBindJSON(&json); err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	userID, _ := c.Get("userID")
	songList, code := songListService.CreateSongList(models.SongList{UserID: userID.(uint), Name: json.Name, Pic: json.Pic, Introduction: json.Introduction})
	appG.Response(code, songList)
}
