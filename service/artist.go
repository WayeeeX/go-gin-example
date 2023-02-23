package service

import (
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
)

type ArtistService struct {
}

func (s *ArtistService) GetSelectList(req request.PageQuery) response.ArtistSelectList {
	return artistModel.GetSelectList(req)
}
