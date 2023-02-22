package models

import "github.com/WayeeeX/go-gin-example/models/common"

type Album struct {
	common.Model
	ArtistID     uint64            `json:"artist_id"`
	Name         string            `json:"name"`
	Introduction string            `json:"introduction"`
	Genre        string            `json:"genre"`
	Publisher    string            `json:"publisher"`
	Status       *int              `json:"status"`
	ReleaseTime  *common.LocalTime `json:"release_time"`
}
