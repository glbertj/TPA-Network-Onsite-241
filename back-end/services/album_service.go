package services

import (
	"back-end/data/request"
	"back-end/data/response"
)

type AlbumService interface {
	GetAlbumsByArtist(artistId string) (res []response.AlbumResponse, err error)
	GetRandomAlbum() (res []response.AlbumResponse, err error)
	CreateAlbum(req request.AlbumRequest) (albumResponse response.AlbumResponse, err error)
}
