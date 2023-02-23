package service

import (
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
)

type AlbumService struct {
}

func (a AlbumService) GetSelectList(req request.PageQuery) response.AlbumSelectList {
	return albumModel.GetSelectList(req)
}
