package response

type SearchResponse struct {
	Payload interface{} `json:"payload"`
	Type    string      `json:"type"`
	Count   int         `json:"count"`
	Title   string      `json:"title"`
}

type SearchResultResponse struct {
	Song  SongResponse `json:"song"`
	Type  string       `json:"type"`
	Count int          `json:"count"`
	Title string       `json:"title"`
}

type ArtistSearch struct {
	ArtistId    string `gorm:"primaryKey"`
	UserId      string
	Username    string
	FollowCount int `gorm:"column:follow_count"`
}

type AlbumSearch struct {
	AlbumId   string
	Title     string
	PlayCount int `gorm:"column:play_count"`
}

type SongSearch struct {
	SongId    string
	Title     string
	PlayCount int `gorm:"column:play_count"`
}
