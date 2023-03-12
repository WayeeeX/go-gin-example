package request

type PlaylistCatList struct {
	PageQuery
	PID uint64 `form:"pid"`
}
type UpdatePlaylistCat struct {
	IdPrimaryKey
	Name  string `json:"name"`
	PID   uint64 `json:"pid" gorm:"column:pid"`
	Order uint64 `json:"order"`
}
type CreatePlaylistCat struct {
	Name  string `json:"name"`
	PID   uint64 `json:"pid"  gorm:"column:pid"`
	Order uint64 `json:"order"`
}
