package services

import (
	"back-end/data/request"
	"back-end/data/response"
)

type PlaylistService interface {
	Create(playlist request.PlayListRequest) error
	GetAll() (res []response.PlayListResponse, err error)
	GetByUserID(id string) (res []response.PlayListResponse, err error)
	GetPlaylistByID(id string) (res response.PlayListResponse, err error)
	CreateDetail(playlistDetail request.PlayListDetailRequest) error
	DeletePlaylistByID(userId string, id string) error
	DeletePlaylistDetailByID(userId string, id string, detailId string) error
}
