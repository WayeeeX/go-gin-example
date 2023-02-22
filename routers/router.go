package routers

import (
	"fmt"
	"github.com/WayeeeX/go-gin-example/middleware"
	"github.com/WayeeeX/go-gin-example/pkg/setting"
	"github.com/WayeeeX/go-gin-example/pkg/upload"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

var Router = &http.Server{}

func InitRouters() {
	routers := GetRouters()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	Router = &http.Server{
		Addr:           endPoint,
		Handler:        routers,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	Router.ListenAndServe()
}

func GetRouters() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorRecovery())
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitAdminRouter(r)
	InitFrontendRouter(r)
	return r
}
