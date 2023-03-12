package models

import (
	"github.com/WayeeeX/go-gin-example/models/common"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
	"github.com/WayeeeX/go-gin-example/pkg/util"
)

type LoginRecord struct {
	ID        uint64            `gorm:"primaryKey" json:"id"`
	UserID    uint64            `json:"user_id"`
	LoginIP   string            `json:"login_ip"`
	LoginTime *common.LocalTime `gorm:"autoCreateTime" json:"login_time"`
}

func (l *LoginRecord) Create(loginRecord LoginRecord) LoginRecord {
	err := DB.Create(&loginRecord).Error
	if err != nil {
		panic(err)
	}
	return loginRecord
}
func (l *LoginRecord) GetList(req request.PageQuery) (res response.LogLoginList) {
	DB.Table("tb_login_record log").Select("log.*,u.id user_id,u.nickname user_name").Joins("left join tb_user u on u.id = log.user_id").Where("u.nickname like ? or log.login_ip like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%").Count(&res.Total).Limit(req.PageSize).Offset(util.GetOffset(req)).Order("login_time desc").Scan(&res.Logs)
	return
}
