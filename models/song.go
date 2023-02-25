package models

import (
	"github.com/WayeeeX/go-gin-example/models/common"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
	"github.com/WayeeeX/go-gin-example/pkg/util"
)

type Song struct {
	common.Model
	ArtistID   uint64 `json:"artist_id"`
	AlbumID    uint64 `json:"album_id"`
	Name       string `json:"name"`
	Status     *int   `json:"status"`
	Url        string `json:"url"`
	IsOriginal *int   `json:"is_original"`
	Pic        string `json:"pic"`
	Duration   int    `json:"duration"`
	Lyric      string `json:"lyrics"`
}

func (s Song) Create(req request.CreateSong) {
	err := db.Table("tb_song").Create(&req).Error
	if err != nil {
		panic(err)
	}
}

func (s Song) Delete(req request.IdsJson) int {
	//TODO implement me
	panic("implement me")
}

func (s Song) Update(song Song) Song {
	//TODO implement me
	panic("implement me")
}

func (s Song) GetOne(song Song) Song {
	//TODO implement me
	panic("implement me")
}
func (s Song) GetList(req request.PageQuery) (res response.SongList) {
	db.Table("tb_song s").Select("s.*,al.name album_name,al.id album_id ,ar.name artist_name,ar.id artist_id").Joins(" join tb_album al on s.album_id = al.id left join tb_artist ar on al.artist_id = ar.id").Count(&res.Total).Where("s.name like ? or ar.name like ? or al.name like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%", "%"+req.Keyword+"%").Limit(req.PageSize).Offset(util.GetOffset(req)).Scan(&res.Songs)
	return
}
