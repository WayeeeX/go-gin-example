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

func GetPlaylistCatList(c *gin.Context) {
	req := app.BindValidQuery[request.PlaylistCatList](c)
	var total uint64
	var categories []response.PlaylistCat
	var parent models.PlaylistCat
	if req.PID == 0 {
		err := models.DB.Table("tb_playlist_category").Where("name like ? and pid = 0", "%"+req.Keyword+"%").Count(&total).Limit(req.PageSize).Offset(util.GetOffset(req.PageQuery)).Order("create_time desc").Scan(&categories).Error
		if err != nil {
			panic(err)
		}
	} else {
		err := models.DB.Table("tb_playlist_category").Where("name like ? and pid = ?", "%"+req.Keyword+"%", req.PID).Count(&total).Limit(req.PageSize).Offset(util.GetOffset(req.PageQuery)).Order("create_time desc").Scan(&categories).Error
		parent = models.GetOne[models.PlaylistCat](parent, "id = ?", req.PID)
		if err != nil {
			panic(err)
		}
	}
	var subTitle string
	pname := parent.Name
	if pname != "" {
		subTitle = pname + "的子分类列表"
	}
	util.Response(c, e.SUCCESS, gin.H{
		"categories": categories,
		"sub_title":  subTitle,
		"total":      total,
	})
	return
}
func UpdatePlaylistCat(c *gin.Context) {
	json := app.BindJson[request.UpdatePlaylistCat](c)
	err := models.DB.Table("tb_playlist_category").Where("id = ?", json.ID).Updates(json).Error
	if err != nil {
		panic(err)
	}
	util.OK(c)
	return
}
func CreatePlaylistCat(c *gin.Context) {
	json := app.BindJson[request.CreatePlaylistCat](c)
	err := models.DB.Table("tb_playlist_category").Create(json).Error
	if err != nil {
		panic(err)
	}
	util.OK(c)
	return
}
func DeletePlaylistCat(c *gin.Context) {
	json := app.BindJson[request.IdsJson](c)
	models.Delete[models.PlaylistCat](models.PlaylistCat{}, "id = (?)", json.Ids)
	util.OK(c)
	return
}
