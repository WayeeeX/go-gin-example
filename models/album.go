package models

type Album struct {
	Model
	ArtistID     uint64     `json:"artist_id"`
	Name         string     `json:"name"`
	Introduction string     `json:"introduction"`
	Genre        string     `json:"genre"`
	Publisher    string     `json:"publisher"`
	Status       *int       `json:"status"`
	ReleaseTime  *LocalTime `json:"release_time"`
}
