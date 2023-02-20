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
