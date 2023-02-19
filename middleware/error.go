package middleware

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover(); err != nil {
			appG := app.GetGin(c)
			// 处理 panic(xxx) 的操作
			if code, ok := err.(int); ok { // panic(code) 根据错误码获取 msg
				appG.Response(code, nil)
			} else if msg, ok := err.(string); ok { // panic(string) 返回 string
				appG.FailMessage(msg)
			}
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			c.Next()
		}
	}
}
