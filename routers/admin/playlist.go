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

func CreatePlaylist(c *gin.Context) {
	json := app.BindJson[request.CreatePlaylist](c)
	playlist := util.CopyProperties[models.Playlist](json)
	models.Create[models.Playlist](&playlist)
	util.OK(c)
	return
}
func DeletePlaylist(c *gin.Context) {
	json := app.BindJson[request.IdsJson](c)
	models.Delete[models.Playlist](models.Playlist{}, "id = (?)", json.Ids)
	util.OK(c)
	return
}
func UpdatePlaylist(c *gin.Context) {
	json := app.BindJson[request.UpdatePlaylist](c)
	playlist := util.CopyProperties[models.Playlist](json)
	models.Updates[models.Playlist](&playlist, "id = ?", playlist.ID)
	util.OK(c)
	return
}
func GetPlaylistList(c *gin.Context) {
	req := app.BindValidQuery[request.PlaylistList](c)
	var total uint64
	var playlists []response.Playlist
	var user models.User
	if req.UserID == 0 {
		err := models.DB.Table("tb_playlist pl").Select("pl.*,u.nickname user_name").Joins("left join tb_user u on u.id = pl.user_id").Where("concat(pl.name,u.nickname) like ?", "%"+req.Keyword+"%").Count(&total).Limit(req.PageSize).Offset(util.GetOffset(req.PageQuery)).Order("create_time desc").Scan(&playlists).Error
		if err != nil {
			panic(err)
		}
	} else {
		err := models.DB.Table("tb_playlist").Where("name like ? and user_id = ?", "%"+req.Keyword+"%", req.UserID).Count(&total).Limit(req.PageSize).Offset(util.GetOffset(req.PageQuery)).Order("create_time desc").Scan(&playlists).Error
		user = models.GetOne[models.User](user, "id = ?", req.UserID)
		for _, v := range playlists {
			v.UserName = user.Nickname
		}
		if err != nil {
			panic(err)
		}
	}
	subTitle := user.Nickname
	util.Response(c, e.SUCCESS, gin.H{
		"playlists": playlists,
		"total":     total,
		"sub_title": subTitle,
	})
	return
}
func GetPlaylistDetail(c *gin.Context) {
	req := app.BindValidQuery[request.IdQuery](c)
	var playlist response.Playlist
	err := models.DB.Table("tb_playlist pl").Select("pl.*,u.nickname user_name").Joins("left join tb_user u on u.id = pl.user_id").Where("pl.id = ?", req.Id).Scan(&playlist).Error
	if err != nil {
		panic(err)
	}
	util.Response(c, e.SUCCESS, gin.H{
		"playlist": playlist,
	})
	return
}
