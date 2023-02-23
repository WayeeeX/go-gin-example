package response

type AlbumSelectList struct {
	Albums []albumSelect `json:"albums"`
	Total  uint64        `json:"total"`
}
type albumSelect struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Pic        string `json:"pic"`
	ArtistName string `json:"artist_name"`
}
