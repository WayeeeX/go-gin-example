package routers

import (
	"github.com/WayeeeX/go-gin-example/middleware"
	v1 "github.com/WayeeeX/go-gin-example/routers/admin"
	"github.com/gin-gonic/gin"
)

func InitAdminRouter(r *gin.Engine) *gin.Engine {
	//无需鉴权路由
	r.POST("/admin/login", v1.AdminLogin)

	admin := r.Group("/admin")
	needAdmin := true
	admin.Use(middleware.AuthJWT(needAdmin))

	artist := admin.Group("/artist")
	{
		artist.GET("/list", v1.GetArtistList)
		artist.GET("/detail", v1.GetArtistDetail)
		artist.POST("/update", v1.UpdateArtist)
		artist.POST("/create", v1.CreateArtist)
		artist.POST("/delete")
		artist.GET("/selectList", v1.GetArtistSelectList)
	}

	user := admin.Group("/user")
	{
		user.GET("/list", v1.GetUserList)
		user.GET("/my/detail", v1.GetMyDetail)
		user.GET("/detail", v1.GetUserDetail)
		user.GET("/selectList", v1.GetUserSelectList)
		user.POST("/delete", v1.DeleteUsers)
		user.POST("/update", v1.UpdateUser)
		user.POST("/updateStatus", v1.UpdateUserStatus)
	}

	song := admin.Group("/song")
	{
		song.GET("/list", v1.GetSonglist)
		song.GET("/detail", v1.GetSongDetail)
		song.POST("/create", v1.CreateSong)
		song.POST("/update", v1.UpdateSong)
		song.POST("/updateStatus", v1.UpdateSongStatus)
		song.POST("/delete", v1.DeleteSongs)
	}

	album := admin.Group("/album")
	{
		album.GET("/list", v1.GetAlbumList)
		album.GET("/detail", v1.GetAlbumDetail)
		album.POST("/create", v1.CreateAlbum)
		album.POST("/update", v1.UpdateAlbum)
		album.POST("/updateStatus")
		album.POST("/delete")
		album.GET("/selectList", v1.GetAlbumSelectList)
	}

	playlist := admin.Group("/playlist")
	{
		playlist.GET("/list", v1.GetPlaylistList)
		playlist.GET("/detail", v1.GetPlaylistDetail)
		playlist.POST("/update", v1.UpdatePlaylist)
		playlist.POST("/create", v1.CreatePlaylist)
		playlist.POST("/delete", v1.DeletePlaylist)
		playlistCategory := playlist.Group("/category")
		{
			playlistCategory.GET("/list", v1.GetPlaylistCatList)
			playlistCategory.GET("/detail")
			playlistCategory.POST("/update", v1.UpdatePlaylistCat)
			playlistCategory.POST("/create", v1.CreatePlaylistCat)
			playlistCategory.POST("/delete", v1.DeletePlaylistCat)
		}
	}

	loginRecord := admin.Group("/log/login")
	{
		loginRecord.GET("/list", v1.GetLoginLogList)
	}

	return r
}
