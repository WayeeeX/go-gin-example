package routers

import (
	"github.com/WayeeeX/go-gin-example/pkg/e"
	"github.com/WayeeeX/go-gin-example/pkg/upload"
	"github.com/WayeeeX/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		util.FailMessage(c, err.Error())
		return
	}

	if image == nil {
		util.Response(c, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetFileName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		util.Response(c, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckFile(fullPath)
	if err != nil {
		util.Response(c, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		util.Response(c, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	util.Response(c, e.SUCCESS, map[string]string{
		"url":      upload.GetImageFullUrl(imageName),
		"save_url": "/" + savePath + imageName,
	})
}

func UploadLyric(c *gin.Context) {
	file, fileInfo, err := c.Request.FormFile("lyric")

	if err != nil {
		util.FailMessage(c, err.Error())
		return
	}

	if fileInfo == nil {
		util.Response(c, e.INVALID_PARAMS, nil)
		return
	}

	fileName := upload.GetFileName(fileInfo.Filename)
	fullPath := upload.GetLyricFullPath()
	savePath := upload.GetLyricPath()
	src := fullPath + fileName

	if !upload.CheckLyricExt(fileName) || !upload.CheckLyricSize(file) {
		util.FailMessage(c, "校验失败,格式或大小有问题")
		return
	}

	err = upload.CheckFile(fullPath)
	if err != nil {
		util.FailMessage(c, "检查文件失败")
		return
	}

	if err := c.SaveUploadedFile(fileInfo, src); err != nil {
		util.FailMessage(c, "保存文件失败")
		return
	}

	util.Response(c, e.SUCCESS, map[string]string{
		"url":      upload.GetLyricFullUrl(fileName),
		"save_url": "/" + savePath + fileName,
	})
}

func UploadMusic(c *gin.Context) {
	file, fileInfo, err := c.Request.FormFile("music")

	if err != nil {
		util.FailMessage(c, err.Error())
		return
	}

	if fileInfo == nil {
		util.Response(c, e.INVALID_PARAMS, nil)
		return
	}
	fileName := upload.GetFileName(fileInfo.Filename)
	fullPath := upload.GetMusicFullPath()
	savePath := upload.GetMusicPath()
	src := fullPath + fileName

	if !upload.CheckMusicExt(fileName) || !upload.CheckMusicSize(file) {
		util.FailMessage(c, "校验失败,格式或大小有问题")
		return
	}

	err = upload.CheckFile(fullPath)
	if err != nil {
		util.FailMessage(c, "检查文件失败")
		return
	}

	if err := c.SaveUploadedFile(fileInfo, src); err != nil {
		util.FailMessage(c, "保存文件失败")
		return
	}

	util.Response(c, e.SUCCESS, map[string]string{
		"url":      upload.GetMusicFullUrl(fileName),
		"save_url": "/" + savePath + fileName,
	})
}
