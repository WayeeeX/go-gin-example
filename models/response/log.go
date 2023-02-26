package response

import (
	"github.com/WayeeeX/go-gin-example/models/common"
)

type LogLoginList struct {
	Logs  []LogLogin `json:"logs"`
	Total uint64     `json:"total"`
}

type LogLogin struct {
	ID        uint64            `json:"id"`
	UserID    uint64            `json:"user_id"`
	UserName  string            `json:"user_name"`
	LoginIP   string            `json:"login_ip"`
	LoginTime *common.LocalTime `json:"login_time"`
}
