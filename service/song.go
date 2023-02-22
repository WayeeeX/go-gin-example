package service

import (
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
)

type SongService struct {
}

func (s SongService) Create(req request.CreateSong) {
	songModel.Create(req)
}
func (s SongService) GetSongList(req request.PageQuery) (res response.SongList) {
	return songModel.GetList(req)
}
