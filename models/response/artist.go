package response

type ArtistSelectList struct {
	Artists []artistSelect `json:"artists"`
	Total   uint64         `json:"total"`
}
type artistSelect struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Pic  string `json:"pic"`
}
