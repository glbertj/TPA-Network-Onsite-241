package response

import "time"

type AlbumResponse struct {
	AlbumId  string         `json:"albumId"`
	Title    string         `json:"title"`
	Type     string         `json:"type"`
	Banner   string         `json:"banner"`
	Release  time.Time      `json:"release"`
	ArtistId string         `json:"artistId"`
	Artist   ArtistResponse `json:"artist" gorm:"foreignKey:ArtistId;references:ArtistId"`
}

func (AlbumResponse) TableName() string {
	return "albums"
}
