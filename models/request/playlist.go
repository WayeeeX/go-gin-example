package request

import "github.com/WayeeeX/go-gin-example/models/common"

type CreatePlaylist struct {
	UserID       uint64 `json:"user_id" binding:"required"`
	Name         string `json:"name" binding:"required,min=4,max=32"`
	Pic          string `json:"pic" binding:"required"`
	Introduction string `json:"introduction" `
}
type UpdatePlaylist struct {
	IdPrimaryKey
	UserID       uint64 `json:"user_id" binding:"required"`
	Name         string `json:"name" binding:"required,min=4,max=32"`
	Pic          string `json:"pic" binding:"required"`
	Introduction string `json:"introduction" `
}
type PlaylistList struct {
	PageQuery
	UserID uint64 `form:"user_id"`
}
type Playlist struct {
	common.Model
	UserID       uint64 `json:"user_id"`
	Name         string `json:"name"`
	Pic          string `json:"pic"`
	Introduction string `json:"introduction"`
}
