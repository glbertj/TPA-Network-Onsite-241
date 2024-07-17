package response

import "time"

type SongResponse struct {
	SongId      string         `json:"songId"`
	Title       string         `json:"title"`
	ArtistId    string         `json:"artistId"`
	AlbumId     string         `json:"albumId"`
	ReleaseDate time.Time      `json:"releaseDate"`
	Duration    int            `json:"duration"`
	File        string         `json:"file"`
	Album       AlbumResponse  `json:"album" gorm:"foreignKey:AlbumId;references:AlbumId"`
	Play        []PlayResponse `json:"play" gorm:"foreignKey:SongId;references:SongId"`
	Artist      ArtistResponse `json:"artist" gorm:"foreignKey:ArtistId;references:ArtistId"`
}

func (SongResponse) TableName() string {
	return "songs"
}
