package admin

import (
	"github.com/WayeeeX/go-gin-example/service"
)

var (
	userService     service.UserService
	songListService service.PlaylistService
	songService     service.SongService
	albumService    service.AlbumService
	artistService   service.ArtistService
	logLoginService service.LogLoginService
)
