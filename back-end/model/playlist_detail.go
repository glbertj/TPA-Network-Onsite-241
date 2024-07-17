package model

import (
	"back-end/data/response"
	"time"
)

type PlaylistDetails struct {
	PlaylistDetailId string `gorm:"primaryKey"`
	PlaylistId       string
	SongId           string
	DateAdded        time.Time                 `gorm:"autoCreateTime"`
	Playlist         response.PlayListResponse `gorm:"foreignKey:PlaylistId;references:PlaylistId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Song             response.SongResponse     `gorm:"foreignKey:SongId;references:SongId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
