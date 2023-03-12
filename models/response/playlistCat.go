package response

import "github.com/WayeeeX/go-gin-example/models/common"

type PlaylistCat struct {
	common.Model
	Name     string `json:"name"`
	PID      uint64 `json:"pid" gorm:"column:pid"`
	Order    uint64 `json:"order"`
	SubTitle string `json:"sub_title"`
}
