package response

import "time"

type PlaylistDetailResponse struct {
	PlaylistDetailId string       `json:"playlistDetailId"`
	PlaylistId       string       `json:"playlistId"`
	SongId           string       `json:"songId"`
	DateAdded        time.Time    `json:"dateAdded"`
	Song             SongResponse `json:"song" gorm:"foreignKey:SongId;references:SongId"`
}

func (PlaylistDetailResponse) TableName() string {
	return "playlist_details"
}
