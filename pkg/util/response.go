package util

import (
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseEntity struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(c *gin.Context, errCode int, data interface{}) {
	if errCode != e.SUCCESS {
		data = nil
	}
	c.JSON(http.StatusOK, ResponseEntity{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
func OK(c *gin.Context) {
	Response(c, e.SUCCESS, nil)
}
func FailMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, ResponseEntity{
		Code: e.ERROR,
		Msg:  message,
		Data: nil,
	})
}
