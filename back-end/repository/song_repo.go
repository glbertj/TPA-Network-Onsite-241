package repository

import (
	"back-end/data/response"
	"back-end/model"
)

type SongRepository interface {
	GetAllSong() (res []model.Song, err error)
	GetSortedSong() (res []model.Song, err error)
	GetSongById(id string) (res model.Song, err error)
	FindSongByTitle(title string) (res []response.SongSearch, err error)
	GetSongByArtist(artistId string) (res []model.Song, err error)
	GetSongByAlbum(albumId string) (res []model.Song, err error)
	CreateSong(song model.Song) error
	GetTop5TrackFromAlbum(albumId string) (res []model.Song, err error)
	GetTop5TrackFromArtist(artistId string) (res []model.Song, err error)
}
