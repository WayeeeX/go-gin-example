package service

import (
	"github.com/WayeeeX/go-gin-example/models"
	"github.com/WayeeeX/go-gin-example/models/request"
)

type SongService struct {
}

func (s SongService) Create(req request.CreateSong) {
	songModel.Create(req)
}
func (s SongService) GetSongList(req request.PageQuery) (songs []models.Song, total int) {
	return songModel.GetList(req)
}
