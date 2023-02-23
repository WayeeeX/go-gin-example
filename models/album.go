package models

import (
	"github.com/WayeeeX/go-gin-example/models/common"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
	"github.com/WayeeeX/go-gin-example/pkg/util"
)

type Album struct {
	common.Model
	ArtistID     uint64            `json:"artist_id"`
	Name         string            `json:"name"`
	Pic          string            `json:"pic"`
	Introduction string            `json:"introduction"`
	Genre        string            `json:"genre"`
	Publisher    string            `json:"publisher"`
	Status       *int              `json:"status"`
	ReleaseTime  *common.LocalTime `json:"release_time"`
}

func (a *Album) GetSelectList(req request.PageQuery) (res response.AlbumSelectList) {
	db.Table("tb_album al").Select("al.id,al.name,al.pic,ar.name artist_name").Joins("left join tb_artist ar on al.artist_id = ar.id").Count(&res.Total).Where("al.name like ? or ar.name like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%").Limit(req.PageSize).Offset(util.GetOffset(req)).Scan(&res.Albums)
	return
}
