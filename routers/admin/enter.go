package admin

import (
	"github.com/WayeeeX/go-gin-example/service"
)

var (
	userService     service.UserService
	songListService service.SongListService
	songService     service.SongService
	albumService    service.AlbumService
	artistService   service.ArtistService
)
