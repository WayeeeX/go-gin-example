package service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
)

type SongListService struct {
}

func (s SongListService) CreateSongList(songList models.SongList) (models.SongList, int) {

	return songListModel.Create(songList), e.SUCCESS
}
