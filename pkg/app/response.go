package app

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/WayeeeX/go-gin-example/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(errCode int, data interface{}) {
	if errCode != e.SUCCESS {
		data = nil
	}
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
func (g *Gin) OK() {
	g.Response(e.SUCCESS, nil)
}
func (g *Gin) FailMessage(message string) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.ERROR,
		Msg:  message,
		Data: nil,
	})
}
func GetGin(c *gin.Context) Gin {
	return Gin{C: c}
}
