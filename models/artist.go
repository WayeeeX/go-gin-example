package models

import (
	"github.com/WayeeeX/go-gin-example/models/common"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
	"github.com/WayeeeX/go-gin-example/pkg/util"
)

type Artist struct {
	common.Model
	Type         int    `json:"type"`
	Region       int    `json:"region"`
	Name         string `json:"name"`
	Pic          string `json:"pic"`
	Introduction string `json:"introduction"`
}

func (a *Artist) GetSelectList(req request.PageQuery) (res response.ArtistSelectList) {
	db.Table("tb_artist ar").Select("ar.id,ar.name,ar.pic").Count(&res.Total).Where("ar.name like ?", "%"+req.Keyword+"%").Limit(req.PageSize).Offset(util.GetOffset(req)).Scan(&res.Artists)
	return
}
