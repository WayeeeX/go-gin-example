package admin

import (
	"github.com/WayeeeX/go-gin-example/models"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/app"
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

type result struct {
	Songs      models.Song
	AlbumName  string
	AlbumPic   string
	ArtistName string
	ArtistPic  string
}

func GetSongList(c *gin.Context) {
	req := app.BindValidQuery[request.PageQuery](c)
	util.Response(c, e.SUCCESS, songService.GetSongList(req))
	return
}
func CreateSong(c *gin.Context) {
	appG := app.GetGin(c)
	song := util.CopyProperties[models.Song](app.BindJson[request.CreateSong](c))
	models.Create(&song)
	appG.OK()
	return
}
func GetSongDetail(c *gin.Context) {
	song := models.GetOne(models.Song{}, "id = ?", app.BindValidQuery[request.IdQuery](c).Id)
	util.Response(c, e.SUCCESS, song)
	return
}
func UpdateSong(c *gin.Context) {
	song := util.CopyProperties[models.Song](app.BindJson[request.UpdateSong](c))
	models.Updates[models.Song](&song, "id = ?", song.ID)
	util.OK(c)
}
func UpdateSongStatus(c *gin.Context) {
	json := app.BindJson[request.UpdateStatus](c)
	models.UpdatesMap[models.Song](&models.Song{}, map[string]any{"status": json.Status}, "id IN (?)", json.Ids)
	util.OK(c)
}

func DeleteSongs(c *gin.Context) {
	json := app.BindJson[request.IdsJson](c)
	models.Delete[models.Song](models.Song{}, "id IN (?)", json.Ids)
	util.OK(c)
}
