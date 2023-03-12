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
func (s SongService) GetSonglist(req request.PageQuery) (res response.SngList) {
	return songModel.GetList(req)
}

func (s SongService) GetSongDetail(id uint64) (res response.Song) {
	return songModel.GetDetail(id)
}
