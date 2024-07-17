package services

import (
	"back-end/data/request"
	"back-end/data/response"
)

type SongService interface {
	GetAllSong() (res []response.SongResponse, err error)
	GetSongById(id string) (res response.SongResponse, err error)
	GetSongByArtist(artistId string) (res []response.SongResponse, err error)
	GetSongByAlbum(albumId string) (res []response.SongResponse, err error)
	CreateSong(song request.SongRequest) error
}
