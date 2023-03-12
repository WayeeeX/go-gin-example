package response

import "github.com/WayeeeX/go-gin-example/models/common"

type Playlist struct {
	common.Model
	UserID       uint   `json:"user_id"`
	Name         string `json:"name"`
	Pic          string `json:"pic"`
	Introduction string `json:"introduction"`
	UserName     string `json:"user_name"`
}
