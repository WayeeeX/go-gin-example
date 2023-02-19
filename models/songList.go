package models

type SongList struct {
	Model
	UserID       uint   `json:"user_id"`
	Name         string `json:"name"`
	Pic          string `json:"pic"`
	Introduction string `json:"introduction"`
}
type songListModel interface {
	Create() SongList
	Update() SongList
	Delete() bool
}

func (s SongList) Create(model SongList) SongList {
	err := db.Create(&model).Error
	if err != nil {
		panic(err)
	}
	return model
}

func (s SongList) Update() SongList {
	//TODO implement me
	panic("implement me")
}

func (s SongList) Delete() bool {
	//TODO implement me
	panic("implement me")
}
