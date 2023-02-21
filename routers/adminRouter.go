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
		artist.GET("/list")
		artist.POST("/update")
		artist.POST("/delete")
	}

	user := admin.Group("/user")
	{
		user.GET("/list", v1.GetUserList)
		user.POST("/delete", v1.DeleteUsers)
		user.POST("/updateStatus", v1.UpdateUserStatus)
	}

	song := admin.Group("/song")
	{
		song.GET("/list", v1.GetSongList)
		song.GET("/detail", v1.GetSongDetail)
		song.POST("/create", v1.CreateSong)
		song.POST("/update", v1.UpdateSong)
		song.POST("/updateStatus", v1.UpdateSongStatus)
		song.POST("/delete", v1.DeleteSongs)
	}

	album := admin.Group("/album")
	{
		album.GET("/list", v1.GetAlbumList)
		album.GET("/detail")
		album.POST("/create", v1.CreateAlbum)
		album.POST("/update")
		album.POST("/updateStatus")
		album.POST("/delete")
	}

	songList := admin.Group("/songlist")
	{
		songList.GET("/list")
		songList.GET("/detail")
		songList.POST("/update")
		songList.POST("/delete")
		songListTag := songList.Group("/tag")
		{
			songListTag.GET("/list")
			songListTag.GET("/detail")
			songListTag.POST("/update")
			songListTag.POST("/delete")
		}
	}

	loginRecord := admin.Group("/record/login")
	{
		loginRecord.GET("/list")
	}

	return r
}
