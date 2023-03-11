package request

type Login struct {
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

type Register struct {
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=32"`
	Nickname string `json:"nickname" binding:"required,min=4,max=32"`
}

type UpdateUserAdmin struct {
	ID       uint64 `json:"id" binding:"required,number,gt=0"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone" binding:"len=11"`
	Avatar   string `json:"avatar"`
	Status   *int   `json:"status"`
	Role     *int   `json:"role"`
}
