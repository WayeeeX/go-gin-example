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

func GetArtistSelectList(c *gin.Context) {
	req := app.BindValidQuery[request.PageQuery](c)
	util.Response(c, e.SUCCESS, artistService.GetSelectList(req))
	return
}
func GetArtistList(c *gin.Context) {
	req := app.BindValidQuery[request.ArtistList](c)
	res := response.ArtistList{}
	keyword := "%" + req.Keyword + "%"
	models.DB.Table("tb_artist s").Where("category like ? or nationality like ? or name like ?", keyword, keyword, keyword).Count(&res.Total).Limit(req.PageSize).Offset(util.GetOffset(req.PageQuery)).Order("create_time desc").Scan(&res.Artists)
	util.Response(c, e.SUCCESS, res)
	return
}
func GetArtistDetail(c *gin.Context) {
	id := app.BindValidQuery[request.IdQuery](c).Id
	util.Response(c, e.SUCCESS, gin.H{
		"artist": models.GetOne[models.Artist](models.Artist{}, "id = ?", id),
	})
	return
}
func UpdateArtist(c *gin.Context) {
	json := app.BindJson[request.UpdateArtist](c)
	err := models.DB.Table("tb_artist").Where("id = ?", json.ID).Updates(json).Error
	if err != nil {
		panic(err)
	}
	util.OK(c)
}

func CreateArtist(c *gin.Context) {
	artist := util.CopyProperties[models.Artist](app.BindJson[request.CreateArtist](c))
	models.Create[models.Artist](&artist)
	util.OK(c)
	return
}
