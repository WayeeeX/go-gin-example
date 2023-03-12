package models

import "github.com/WayeeeX/go-gin-example/models/common"

type Playlist struct {
	common.Model
	UserID       uint   `json:"user_id"`
	Name         string `json:"name"`
	Pic          string `json:"pic"`
	Introduction string `json:"introduction"`
}
type playlistModel interface {
	Create() Playlist
	Update() Playlist
	Delete() bool
}

func (s Playlist) Create(model Playlist) Playlist {
	err := DB.Create(&model).Error
	if err != nil {
		panic(err)
	}
	return model
}

func (s Playlist) Update() Playlist {
	//TODO implement me
	panic("implement me")
}

func (s Playlist) Delete() bool {
	//TODO implement me
	panic("implement me")
}
