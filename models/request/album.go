package request

import (
	"github.com/WayeeeX/go-gin-example/models/common"
)

type CreateAlbum struct {
	ArtistID     uint64      `json:"artist_id" binding:"required,number"`
	Name         string      `json:"name" binding:"required"`
	Introduction string      `json:"introduction"`
	Genre        string      `json:"genre" binding:"required"`
	Publisher    string      `json:"publisher"`
	ReleaseTime  common.Date `json:"release_time" binding:"required"`
	Pic          string      `json:"pic" binding:"required"`
}
type UpdateAlbum struct {
	IdPrimaryKey
	ArtistID     uint64      `json:"artist_id" binding:"required,number"`
	Name         string      `json:"name" binding:"required"`
	Introduction string      `json:"introduction"`
	Genre        string      `json:"genre" binding:"required"`
	Publisher    string      `json:"publisher"`
	ReleaseTime  common.Date `json:"release_time" binding:"required"`
	Pic          string      `json:"pic" binding:"required"`
}
type AlbumList struct {
	PageQuery
	ArtistID uint64 `form:"artist_id" binding:"number"`
}
