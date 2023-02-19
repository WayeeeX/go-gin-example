package util

import (
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type Page struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

// GetPage get page parameters
func GetPage(c *gin.Context) (page Page) {
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()
	if pageNum == 0 {
		page.PageNum = 1
	}
	if pageSize == 0 {
		page.PageSize = setting.AppSetting.PageSize
	}
	return page
}
func GetOffset(page Page) (offset int) {
	return (page.PageNum - 1) * page.PageSize
}
