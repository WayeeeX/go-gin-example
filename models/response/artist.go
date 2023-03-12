package response

import (
	"github.com/WayeeeX/go-gin-example/models/common"
)

type ArtistSelectList struct {
	Artists []artistSelect `json:"artists"`
	Total   uint64         `json:"total"`
}
type artistSelect struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Pic  string `json:"pic"`
}
type ArtistList struct {
	Artists []artist `json:"artists"`
	Total   uint64   `json:"total"`
}
type artist struct {
	common.Model
	Category     string      `json:"category"`
	Nationality  string      `json:"nationality"`
	Birthday     common.Date `json:"birthday"`
	Name         string      `json:"name"`
	Pic          string      `json:"pic"`
	Introduction string      `json:"introduction"`
}
