package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"github.com/WayeeeX/go-gin-example/models"
	"github.com/WayeeeX/go-gin-example/pkg/gredis"
	"github.com/WayeeeX/go-gin-example/pkg/logging"
	"github.com/WayeeeX/go-gin-example/pkg/setting"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"github.com/WayeeeX/go-gin-example/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/WayeeeX/go-gin-example
// @license.name MIT
// @license.url https://github.com/WayeeeX/go-gin-example/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routers.InitRouters()
	log.Printf("[info] start http server listening %s", setting.ServerSetting.HttpPort)

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
