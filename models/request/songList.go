package request

type CreateSongList struct {
	Name         string `json:"name" binding:"required,min=4,max=32"`
	Pic          string `json:"pic" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
}
