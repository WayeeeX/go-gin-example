package routers

import (
	"github.com/EDDYCJY/go-gin-example/middleware"
	v1 "github.com/EDDYCJY/go-gin-example/routers/admin"
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
		user.POST("/delete")
		user.POST("/update")
		user.POST("/updateStatus")
	}

	song := admin.Group("/song")
	{
		song.GET("/list")
		song.GET("/detail")
		song.POST("/update")
		song.POST("/updateStatus")
		song.POST("/delete")
	}

	album := admin.Group("/album")
	{
		album.GET("/list")
		album.GET("/detail")
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
