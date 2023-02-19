package response

import (
	"github.com/EDDYCJY/go-gin-example/models"
)

type AdminLogin struct {
	User  models.User
	Token string `json:"token"`
}
type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
	Role     int    `json:"role"`
}

type Login struct {
	User  UserInfo `json:"user"`
	Token string   `json:"token"`
}
