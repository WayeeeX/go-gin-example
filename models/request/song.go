package request

type CreateSong struct {
	ArtistID   uint64 `json:"artist_id" binding:"required,number,gt=0"`
	AlbumID    uint64 `json:"album_id" binding:"required,number,gt=0"`
	Name       string `json:"name" binding:"required,max=20,min=3"`
	Status     *int   `json:"status" binding:"required,number,oneof=1 0"`
	Url        string `json:"url" binding:"required,url"`
	IsOriginal *int   `json:"is_original" binding:"required,oneof=1 0"`
	Pic        string `json:"pic" binding:"required"`
	Duration   int    `json:"duration" binding:"required,number"`
	Lyric      string `json:"lyric"`
}

type UpdateSong struct {
	ID         uint64 `json:"id" binding:"required,number,gt=0"`
	ArtistID   uint64 `json:"artist_id" binding:"number,gt=0"`
	AlbumID    uint64 `json:"album_id" binding:"number,gt=0"`
	Name       string `json:"name" binding:"max=20,min=3"`
	Status     *int   `json:"status" binding:"number,oneof=1 0"`
	Url        string `json:"url" binding:"url"`
	IsOriginal *int   `json:"is_original" binding:"oneof=1 0"`
	Pic        string `json:"pic" binding:""`
	Duration   int    `json:"duration" binding:"number"`
	Lyric      string `json:"lyric"`
}
