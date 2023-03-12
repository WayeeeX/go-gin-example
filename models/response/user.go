package response

/*type AdminLogin struct {
	common.Model
	Username      string            `json:"username"`
	Nickname      string            `json:"nickname"`
	Phone         string            `json:"phone"`
	Avatar        string            `json:"avatar"`
	Status        *int              `gorm:"default:1" json:"status"`
	Role          *int              `json:"role"`
	LastLoginTime *common.LocalTime `json:"last_login_time"`
	LastLoginIP   string            `json:"last_login_ip"`
	Token         string            `json:"token"`
}*/
type AdminLogin struct {
	User  UserInfo `json:"user"`
	Token string   `json:"token"`
}
type UserInfo struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Status   *int   `json:"status"`
	Role     *int   `json:"role"`
}

type Login struct {
	User  UserInfo `json:"user"`
	Token string   `json:"token"`
}

type UserSelect struct {
	ID       uint64 `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
