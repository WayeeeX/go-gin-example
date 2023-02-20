package models

import (
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/pkg/util"
)

type Song struct {
	Model
	ArtistID   uint64 `json:"artist_id"`
	AlbumID    uint64 `json:"album_id"`
	Name       string `json:"name"`
	Status     *int   `json:"status"`
	Url        string `json:"url"`
	IsOriginal *int   `json:"is_original"`
	Pic        string `json:"pic"`
	Duration   int    `json:"duration"`
	Lyric      string `json:"lyric"`
}

func (s Song) Create(req request.CreateSong) {
	err := db.Table("tb_song").Create(&req).Error
	if err != nil {
		panic(err)
	}
}

func (s Song) Delete(req request.IdsJson) int {
	//TODO implement me
	panic("implement me")
}

func (s Song) Update(song Song) Song {
	//TODO implement me
	panic("implement me")
}

func (s Song) GetOne(song Song) Song {
	//TODO implement me
	panic("implement me")
}
func (s Song) GetList(req request.PageQuery) (songs []Song, total int) {
	db.Table("tb_song").Count(&total).Limit(req.PageSize).Offset(util.GetOffset(req)).
		Find(&songs)
	return songs, total
}
