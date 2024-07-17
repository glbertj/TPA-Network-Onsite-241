package repository

import (
	"back-end/database"
	"back-end/model"
	"back-end/utils"
	"encoding/json"
	"gorm.io/gorm"
)

type PlaylistRepositoryImpl struct {
	DB  *gorm.DB
	rdb *database.Redis
}

func NewPlaylistRepositoryImpl(DB *gorm.DB, rdb *database.Redis) *PlaylistRepositoryImpl {
	return &PlaylistRepositoryImpl{DB: DB, rdb: rdb}
}

func (p PlaylistRepositoryImpl) GetAll() (res []model.Playlist, err error) {
	panic("implement me")
}

func (p PlaylistRepositoryImpl) GetByUserID(id string) (res []model.Playlist, err error) {
	songs, err := p.rdb.Get(utils.PlaylistKey + id)
	if err != nil {
		err = p.DB.Where("user_id", id).Preload("User").Preload("PlaylistDetails").Preload("PlaylistDetails.Song").Preload("PlaylistDetails.Song.Artist").Preload("PlaylistDetails.Song.Album").Preload("PlaylistDetails.Song.Artist.User").Find(&res).Error
		if err != nil {
			return res, err
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			return res, err
		}
		_ = p.rdb.Set(utils.PlaylistKey+id, string(resJSON))

		return res, nil
	} else {
		if err := json.Unmarshal([]byte(songs), &res); err != nil {
			return res, err
		}
		return res, nil
	}
}

func (p PlaylistRepositoryImpl) GetPlaylistByID(id string) (res model.Playlist, err error) {

	songs, err := p.rdb.Get(utils.PlaylistDetailKey + id)
	if err != nil {
		err = p.DB.Where("playlist_id", id).Preload("User").Preload("PlaylistDetails").Preload("PlaylistDetails.Song").Preload("PlaylistDetails.Song.Artist").Preload("PlaylistDetails.Song.Artist.User").Preload("PlaylistDetails.Song.Album").Find(&res).Error

		if err != nil {
			return res, err
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			return res, err
		}
		_ = p.rdb.Set(utils.PlaylistDetailKey+id, string(resJSON))

		return res, nil
	} else {
		if err := json.Unmarshal([]byte(songs), &res); err != nil {
			return res, err
		}
		return res, nil
	}

}

func (p PlaylistRepositoryImpl) Create(playlist model.Playlist) error {
	err := p.rdb.Del(utils.PlaylistDetailKey + playlist.PlaylistId)
	if err != nil {
		return err
	}
	err = p.rdb.Del(utils.PlaylistKey + playlist.UserId)
	if err != nil {
		return err
	}
	err = p.DB.Create(&playlist).Error
	return err
}

func (p PlaylistRepositoryImpl) CreateDetail(userId string, playlistDetail model.PlaylistDetails) error {
	err := p.rdb.Del(utils.PlaylistDetailKey + playlistDetail.PlaylistId)
	if err != nil {
		return err
	}
	err = p.rdb.Del(utils.PlaylistKey + userId)
	if err != nil {
		return err
	}
	err = p.DB.Create(&playlistDetail).Error
	return err
}

func (p PlaylistRepositoryImpl) DeletePlaylistDetailByID(userId string, id string, detailId string) error {
	err := p.rdb.Del(utils.PlaylistDetailKey + detailId)
	if err != nil {
		return err
	}
	err = p.rdb.Del(utils.PlaylistKey + userId)
	if err != nil {
		return err
	}

	err = p.DB.Where("playlist_detail_id", id).Delete(&model.PlaylistDetails{}).Error
	return err
}

func (p PlaylistRepositoryImpl) DeletePlaylistByID(userId string, id string) error {
	err := p.rdb.Del(utils.PlaylistDetailKey + id)
	if err != nil {
		return err
	}
	err = p.rdb.Del(utils.PlaylistKey + userId)
	if err != nil {
		return err
	}
	err = p.DB.Where("playlist_id", id).Delete(&model.Playlist{}).Error
	return err
}
