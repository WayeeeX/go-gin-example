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
	models.DB.Table("tb_album").Create(json)
	util.OK(c)
	return
}
func UpdateAlbum(c *gin.Context) {
	json := app.BindJson[request.UpdateAlbum](c)
	err := models.DB.Table("tb_album").Where("id = ?", json.ID).Updates(json).Error
	if err != nil {
		panic(err)
	}
	util.OK(c)
	return
}
func GetAlbumList(c *gin.Context) {
	req := app.BindValidQuery[request.AlbumList](c)
	var total uint64
	var albums []response.Album
	searchQuery := "concat(al.name,ar.name,al.genre,al.publisher) like ?"
	if req.ArtistID == 0 {
		err := models.DB.Table("tb_album al").Select("al.*,ar.name artist_name").Joins("left join tb_artist ar on al.artist_id = ar.id").Where(searchQuery, "%"+req.Keyword+"%").Count(&total).Limit(req.PageSize).Offset(util.GetOffset(req.PageQuery)).Order("create_time desc").Scan(&albums).Error
		if err != nil {
			panic(err)
		}
	} else {
		err := models.DB.Table("tb_album al").Select("al.*,ar.name artist_name").Joins("left join tb_artist ar on al.artist_id = ar.id").Where(searchQuery+" and ar.id = ?", "%"+req.Keyword+"%", req.ArtistID).Count(&total).Limit(req.PageSize).Offset(util.GetOffset(req.PageQuery)).Order("create_time desc").Scan(&albums).Error
		if err != nil {
			panic(err)
		}
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

func GetAlbumDetail(c *gin.Context) {
	id := app.BindValidQuery[request.IdQuery](c).Id
	var res response.Album
	err := models.DB.Table("tb_album al").Select("al.*,ar.name artist_name").Joins("left join tb_artist ar on al.artist_id = ar.id").Where("al.id = ?", id).Scan(&res).Error
	if err != nil {
		panic(err)
	}
	util.Response(c, e.SUCCESS, gin.H{
		"album": res,
	})
	return
}
