package models

import "github.com/WayeeeX/go-gin-example/models/common"

type PlaylistCat struct {
	common.Model
	Name  string `json:"name"`
	PID   uint64 `json:"pid"`
	Order uint64 `json:"order"`
}

func (PlaylistCat) TableName() string {
	return "tb_playlist_category"
}
