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

func CreateAlbum(c *gin.Context) {
	json := app.BindJson[request.CreateAlbum](c)
	album := util.CopyProperties[models.Album](json)
	//album.ReleaseTime = models.LocalTime.Value(json.ReleaseTime)
	models.Create[models.Album](&album)
	util.OK(c)
	return
}
func GetAlbumList(c *gin.Context) {
	req := app.BindValidQuery[request.AlbumList](c)
	var total uint64
	var albums []response.Album
	searchQuery := "concat(al.name,ar.name,al.genre,al.publisher) like ?"
	if req.ArtistID == 0 {
		models.DB.Table("tb_album al").Select("al.*,ar.name artist_name").Joins("left join tb_artist ar on al.artist_id = ar.id").Count(&total).Where(searchQuery, "%"+req.Keyword+"%").Limit(req.PageSize).Offset(util.GetOffset(req.PageQuery)).Scan(&albums)
	} else {
		models.DB.Table("tb_album al").Select("al.*,ar.name artist_name").Joins("left join tb_artist ar on al.artist_id = ar.id").Count(&total).Where(searchQuery+" and ar.id = ?", "%"+req.Keyword+"%", req.ArtistID).Limit(req.PageSize).Offset(util.GetOffset(req.PageQuery)).Scan(&albums)
	}

	util.Response(c, e.SUCCESS, gin.H{
		"albums": albums,
		"total":  total,
	})
	return
}

func GetAlbumSelectList(c *gin.Context) {
	req := app.BindValidQuery[request.PageQuery](c)
	util.Response(c, e.SUCCESS, albumService.GetSelectList(req))
	return
}
