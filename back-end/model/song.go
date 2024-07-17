package model

import (
	"back-end/data/response"
	"time"
)

type Song struct {
	SongId      string `gorm:"primaryKey"`
	Title       string
	ArtistId    string
	AlbumId     string
	ReleaseDate time.Time
	Duration    int
	File        string
	Play        []response.PlayResponse `gorm:"foreignKey:SongId;references:SongId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Album       response.AlbumResponse  `gorm:"foreignKey:AlbumId;references:AlbumId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Artist      response.ArtistResponse `gorm:"foreignKey:ArtistId;references:ArtistId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ByPlayCount []Song

func (a ByPlayCount) Len() int           { return len(a) }
func (a ByPlayCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPlayCount) Less(i, j int) bool { return len(a[i].Play) < len(a[j].Play) }
