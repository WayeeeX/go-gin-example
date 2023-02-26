package middleware

import (
	"github.com/WayeeeX/go-gin-example/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 处理 panic(xxx) 的操作
				appG := app.GetGin(c)
				if code, ok := err.(int); ok { // panic(code) 根据错误码获取 msg
					appG.Response(code, nil)
				} else if msg, ok := err.(string); ok { // panic(string) 返回 string
					appG.FailMessage(msg)
				} else if msgMap, ok := err.(validator.ValidationErrorsTranslations); ok {
					msg := ""
					i := 0
					len := len(msgMap)
					for _, v := range msgMap {
						i++
						if i == len {
							msg += v
						} else {
							msg += v + ";"
						}
					}
					appG.FailMessage(msg)
				} else if e, ok := err.(error); ok { // panic(error) 发送消息
					appG.FailMessage(e.Error())
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
