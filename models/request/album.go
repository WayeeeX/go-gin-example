package request

type CreateAlbum struct {
	ArtistID     uint64 `json:"artist_id" binding:"required,number"`
	Name         string `json:"name" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Genre        string `json:"genre" binding:"required"`
	Publisher    string `json:"publisher"`
	Status       *int   `json:"status" binding:"required,oneof=0 1"`
	ReleaseTime  string `json:"release_time"`
}
type AlbumList struct {
	PageQuery
	ArtistID uint64 `form:"artist_id" binding:"number"`
}
