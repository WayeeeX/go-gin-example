package models

import (
	"github.com/WayeeeX/go-gin-example/models/common"
	"github.com/WayeeeX/go-gin-example/models/request"
	"github.com/WayeeeX/go-gin-example/models/response"
	"github.com/WayeeeX/go-gin-example/pkg/util"
)

type Song struct {
	common.Model
	AlbumID    uint64 `json:"album_id"`
	Name       string `json:"name"`
	Status     *int   `json:"status"`
	Url        string `json:"url"`
	IsOriginal *int   `json:"is_original"`
	Pic        string `json:"pic"`
	Duration   int    `json:"duration"`
	Lyric      string `json:"lyric"`
}

func (s Song) Create(req request.CreateSong) {
	err := DB.Table("tb_song").Create(&req).Error
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
func (s Song) GetList(req request.PageQuery) (res response.SngList) {
	DB.Table("tb_song s").Select("s.*,al.name album_name,al.id album_id ,ar.name artist_name,ar.id artist_id").Joins("left join tb_album al on s.album_id = al.id left join tb_artist ar on al.artist_id = ar.id").Where("s.name like ? or ar.name like ? or al.name like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%", "%"+req.Keyword+"%").Count(&res.Total).Limit(req.PageSize).Offset(util.GetOffset(req)).Order("create_time desc").Scan(&res.Songs)
	return
}
func (s Song) GetDetail(id uint64) (res response.Song) {
	song := GetOne(Song{}, "id = ?", id)
	res = util.CopyProperties[response.Song](song)
	DB.Table("tb_album album").Select("album.id album_id,album.name album_name,artist.name artist_name,artist.id artist_id").Joins("left join tb_artist artist on  album.artist_id = artist.id").Where("album.id = ?", song.AlbumID).Scan(&res)
	return
}
