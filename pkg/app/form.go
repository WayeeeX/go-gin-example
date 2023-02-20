package app

import (
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

func BindJson[T any](c *gin.Context) (data T) {
	if err := c.ShouldBindJSON(&data); err != nil {
		panic(e.INVALID_PARAMS)
	}
	return
}
func BindValidQuery[T any](c *gin.Context) (data T) {
	// Query 绑定
	if err := c.ShouldBindQuery(&data); err != nil {
		panic(e.INVALID_PARAMS)
	}
	return data
}
