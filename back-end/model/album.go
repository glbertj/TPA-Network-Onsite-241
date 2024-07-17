package model

import (
	"back-end/data/response"
	"time"
)

type Album struct {
	AlbumId  string `gorm:"primaryKey"`
	ArtistId string
	Title    string
	Type     string
	Banner   string
	Release  time.Time
	Artist   response.ArtistResponse `gorm:"foreignKey:ArtistId;references:ArtistId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
