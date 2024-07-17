package request

import (
	"time"
)

type SongRequest struct {
	SongId      string    `json:"songId" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	ArtistId    string    `json:"artistId" validate:"required"`
	AlbumId     string    `json:"albumId" validate:"required"`
	ReleaseDate time.Time `json:"releaseDate" validate:"required"`
	Duration    int       `json:"duration" validate:"required"`
	File        string    `json:"file" validate:"required"`
}
