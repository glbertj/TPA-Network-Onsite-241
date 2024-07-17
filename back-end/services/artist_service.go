package services

import (
	"back-end/data/request"
	"back-end/data/response"
)

type ArtistService interface {
	GetArtistByUserId(userId string) (res response.ArtistResponse, err error)
	GetArtistByArtistId(artistId string) (res response.ArtistResponse, err error)
	CreateArtist(artist request.ArtistRequest) error
	UpdateVerifyArtist(artistId string) error
	GetUnverifiedArtist() (res []response.ArtistResponse, err error)
	GetUnverifiedArtistByArtistId(artistId string) (res response.ArtistResponse, err error)
	DeleteArtist(userId string, artistId string) error
}
