package repository

import (
	"back-end/data/response"
	"back-end/model"
)

type AlbumRepository interface {
	GetAlbumsByTitle(title string) ([]response.AlbumSearch, error)
	GetAlbumsByArtist(artistId string) ([]model.Album, error)
	GetRandomAlbum() ([]model.Album, error)
	CreateAlbum(album model.Album) error
}
