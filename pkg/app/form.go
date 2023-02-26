package app

import (
	"github.com/WayeeeX/go-gin-example/pkg/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindJson[T any](c *gin.Context) (data T) {
	if err := c.ShouldBindJSON(&data); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			panic(err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		panic(errs.Translate(translator.Trans))
		return
	}
	return
}
func BindValidQuery[T any](c *gin.Context) (data T) {
	// Query 绑定
	if err := c.ShouldBindQuery(&data); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			panic(err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		panic(errs.Translate(translator.Trans))
		return
	}
	return data
}
