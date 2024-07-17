package repository

import "back-end/model"

type PlaylistRepository interface {
	Create(playlist model.Playlist) error
	CreateDetail(userId string, playlistDetail model.PlaylistDetails) error
	GetAll() (res []model.Playlist, err error)
	GetByUserID(id string) (res []model.Playlist, err error)
	GetPlaylistByID(id string) (res model.Playlist, err error)
	DeletePlaylistDetailByID(userId string, id string, detailId string) error
	DeletePlaylistByID(userId string, id string) error
}
