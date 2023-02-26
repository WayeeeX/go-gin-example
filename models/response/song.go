package response

import (
	"github.com/WayeeeX/go-gin-example/models/common"
)

type SongList struct {
	Songs []songListSong `json:"songs"`
	Total uint64         `json:"total"`
}
type songListSong struct {
	common.Model
	Name       string `json:"name"`
	Status     *int   `json:"status"`
	Url        string `json:"url"`
	IsOriginal *int   `json:"is_original"`
	Pic        string `json:"pic"`
	Duration   int    `json:"duration"`
	Lyric      string `json:"lyric"`
	AlbumID    uint64 `json:"album_id"`
	AlbumName  string `json:"album_name"`
	ArtistID   uint64 `json:"artist_id"`
	ArtistName string `json:"artist_name"`
}

type Song struct {
	common.Model
	Name       string `json:"name"`
	Status     *int   `json:"status"`
	Url        string `json:"url"`
	IsOriginal *int   `json:"is_original"`
	Pic        string `json:"pic"`
	Duration   int    `json:"duration"`
	Lyric      string `json:"lyric"`
	AlbumID    uint64 `json:"album_id"`
	AlbumName  string `json:"album_name"`
	ArtistID   uint64 `json:"artist_id"`
	ArtistName string `json:"artist_name"`
}
