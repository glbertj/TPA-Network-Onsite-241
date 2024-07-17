package repository

import (
	"back-end/data/response"
	"back-end/database"
	"back-end/model"
	"back-end/utils"
	"encoding/json"
	"gorm.io/gorm"
)

type AlbumRepositoryImpl struct {
	DB  *gorm.DB
	rdb *database.Redis
}

func NewAlbumRepositoryImpl(DB *gorm.DB, rdb *database.Redis) *AlbumRepositoryImpl {
	return &AlbumRepositoryImpl{DB: DB, rdb: rdb}
}

func (a AlbumRepositoryImpl) GetAlbumsByTitle(title string) (res []response.AlbumSearch, err error) {
	err = a.DB.Table("songs AS s").
		Select("s.album_id,a.title, COUNT(p.play_id) AS play_count").
		Joins("LEFT JOIN plays p ON s.song_id = p.song_id").
		Joins("JOIN albums a ON s.album_id = a.album_id").
		Where("UPPER(a.title) LIKE ?", "%"+title+"%").
		Group("s.album_id,a.title").
		Order("play_count DESC").
		Limit(6).
		Scan(&res).Error
	return
}

func (a AlbumRepositoryImpl) GetAlbumsByArtist(artistId string) (res []model.Album, err error) {
	album, err := a.rdb.Get(utils.AlbumKey + artistId)
	if err != nil {
		err = a.DB.Where("artist_id = ?", artistId).Preload("Artist").Find(&res).Error

		if err != nil {
			return res, err
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			return res, err
		}
		_ = a.rdb.Set(utils.AlbumKey+artistId, string(resJSON))

		return res, nil
	} else {
		if err := json.Unmarshal([]byte(album), &res); err != nil {
			return res, err
		}
		return res, nil
	}
}

func (a AlbumRepositoryImpl) GetRandomAlbum() (res []model.Album, err error) {
	err = a.DB.Order("RANDOM()").Limit(10).Find(&res).Error
	return
}

func (a AlbumRepositoryImpl) CreateAlbum(album model.Album) error {
	err := a.rdb.Del(utils.AlbumKey + album.ArtistId)
	if err != nil {
		return err
	}
	err = a.DB.Create(&album).Error
	return err
}
