package service

import (
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
)

type LogLoginService struct {
}

func (s LogLoginService) GetList(req request.PageQuery) response.LogLoginList {
	return loginRecordModel.GetList(req)
}
