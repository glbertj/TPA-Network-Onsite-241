package repository

import "back-end/model"

type PlayRepository interface {
	Create(play model.Play) error
	Get8LastPlayedSongByUser(userId string) (res []model.Play, err error)
	GetLastPlayedSongByUser(userId string) (res []model.Play, err error)
}
