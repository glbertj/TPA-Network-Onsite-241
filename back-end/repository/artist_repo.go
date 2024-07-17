package repository

import (
	"back-end/data/response"
	"back-end/model"
	"time"
)

type ArtistRepository interface {
	GetArtistByUserId(userId string, verified bool) (res model.Artist, err error)
	GetArtistByArtistId(artistId string, verified bool) (res model.Artist, err error)
	CreateArtist(artist model.Artist) error
	UpdateVerifyArtist(userId string, artistId string, verifiedAt time.Time) error
	GetUnverifiedArtist() (res []model.Artist, err error)
	DeleteArtist(userId string, artistId string) error
	GetArtistByName(name string) (res []response.ArtistSearch, err error)
}
