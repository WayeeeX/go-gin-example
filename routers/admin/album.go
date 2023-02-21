package admin

import (
	"github.com/WayeeeX/go-gin-example/models"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/app"
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

func CreateAlbum(c *gin.Context) {
	json := app.BindJson[request.CreateAlbum](c)
	album := util.CopyProperties[models.Album](json)
	//album.ReleaseTime = models.LocalTime.Value(json.ReleaseTime)
	models.Create[models.Album](&album)
	util.OK(c)
	return
}
func GetAlbumList(c *gin.Context) {
	req := app.BindValidQuery[request.PageQuery](c)
	albums, total := models.List([]models.Album{}, req, "concat(name,introduction,publisher) like ?", "%"+req.Keyword+"%")
	util.Response(c, e.SUCCESS, gin.H{
		"albums": albums,
		"total":  total,
	})
	return
}
