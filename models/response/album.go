package response

import (
	"github.com/WayeeeX/go-gin-example/models/common"
)

type AlbumSelectList struct {
	Albums []albumSelect `json:"albums"`
	Total  uint64        `json:"total"`
}
type albumSelect struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Pic        string `json:"pic"`
	ArtistName string `json:"artist_name"`
}
type Album struct {
	common.Model
	ArtistID     uint64      `json:"artist_id"`
	Name         string      `json:"name"`
	Pic          string      `json:"pic"`
	Introduction string      `json:"introduction"`
	Genre        string      `json:"genre"`
	Publisher    string      `json:"publisher"`
	ReleaseTime  common.Date `json:"release_time"`
	ArtistName   string      `json:"artist_name"`
}
