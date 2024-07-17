package request

import (
	"back-end/data/response"
	"time"
)

type QueueRequest struct {
	SongId      string                  `json:"songId" validate:"required"`
	Title       string                  `json:"title" validate:"required"`
	ArtistId    string                  `json:"artistId" validate:"required"`
	AlbumId     string                  `json:"albumId" validate:"required"`
	ReleaseDate time.Time               `json:"releaseDate" validate:"required"`
	Duration    int                     `json:"duration" validate:"required"`
	File        string                  `json:"file" validate:"required"`
	Play        []response.PlayResponse `json:"play" validate:"required"`
	Album       response.AlbumResponse  `json:"album" validate:"required"`
	Artist      response.ArtistResponse `json:"artist" validate:"required"`
}
