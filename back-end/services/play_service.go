package services

import (
	"back-end/data/request"
	"back-end/data/response"
)

type PlayService interface {
	Create(play request.PlayRequest) error
	Get8LastPlayedSongByUser(userId string) (res []response.PlayResponse, err error)
	GetLastPlayedSongByUser(userId string) (res []response.PlayResponse, err error)
}
