package models

import (
	"github.com/WayeeeX/go-gin-example/models/common"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
	"github.com/WayeeeX/go-gin-example/pkg/util"
)

type Artist struct {
	common.Model
	Category     string      `json:"category"`
	Nationality  string      `json:"nationality"`
	Birthday     common.Date `json:"birthday" gorm:"default:nil",type:date`
	Name         string      `json:"name"`
	Pic          string      `json:"pic"`
	Introduction string      `json:"introduction"`
}

func (a *Artist) GetSelectList(req request.PageQuery) (res response.ArtistSelectList) {
	DB.Table("tb_artist ar").Select("ar.id,ar.name,ar.pic").Where("ar.name like ?", "%"+req.Keyword+"%").Count(&res.Total).Limit(req.PageSize).Offset(util.GetOffset(req)).Scan(&res.Artists)
	return
}
