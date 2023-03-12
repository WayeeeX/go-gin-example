package service

import (
	"github.com/WayeeeX/go-gin-example/models"
	"github.com/WayeeeX/go-gin-example/pkg/e"
)

type PlaylistService struct {
}

func (s PlaylistService) CreatePlaylist(songList models.Playlist) (models.Playlist, int) {

	return songListModel.Create(songList), e.SUCCESS
}
