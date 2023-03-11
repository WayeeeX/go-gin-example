package request

import "github.com/WayeeeX/go-gin-example/models/common"

type ArtistList struct {
	PageQuery
}
type CreateArtist struct {
	Category     string            `json:"category" binding:"required"`
	Nationality  string            `json:"nationality" binding:"required"`
	Birthday     *common.LocalDate `json:"birthday"`
	Name         string            `json:"name" binding:"required"`
	Pic          string            `json:"pic" binding:"required"`
	Introduction string            `json:"introduction"`
}
type UpdateArtist struct {
	ID           uint64           `json:"id" binding:"required,number,gt=0"`
	Category     string           `json:"category" binding:"required"`
	Nationality  string           `json:"nationality" binding:"required"`
	Birthday     common.LocalDate `json:"birthday"`
	Name         string           `json:"name" binding:"required"`
	Pic          string           `json:"pic" binding:"required"`
	Introduction string           `json:"introduction"`
}
