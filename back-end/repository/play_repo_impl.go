package repository

import (
	"back-end/database"
	"back-end/model"
	"gorm.io/gorm"
)

type PlayRepositoryImpl struct {
	DB  *gorm.DB
	rdb *database.Redis
}

func NewPlayRepositoryImpl(DB *gorm.DB, rdb *database.Redis) *PlayRepositoryImpl {
	return &PlayRepositoryImpl{DB: DB, rdb: rdb}
}

func (p PlayRepositoryImpl) Create(play model.Play) error {
	return p.DB.Create(&play).Error
}

func (p PlayRepositoryImpl) Get8LastPlayedSongByUser(userId string) (res []model.Play, err error) {
	subquery := p.DB.
		Table("plays").
		Select("song_id, MAX(played_at) as max_played_at").
		Where("user_id = ?", userId).
		Group("song_id")

	err = p.DB.
		Table("plays AS p").
		Joins("JOIN (?) AS sub ON p.song_id = sub.song_id AND p.played_at = sub.max_played_at", subquery).
		Where("p.user_id = ?", userId).
		Order("p.played_at DESC").
		Limit(8).
		Preload("Song").
		Preload("Song.Artist").
		Preload("Song.Album").
		Find(&res).Error
	return
}

func (p PlayRepositoryImpl) GetLastPlayedSongByUser(userId string) (res []model.Play, err error) {
	//err = p.DB.Where("user_id = ?", userId).Order("played_at DESC").Limit(8).Preload("Song").Preload("Song.Artist").Preload("Song.Album").Find(&res).Error
	subquery := p.DB.
		Table("plays").
		Select("song_id, MAX(played_at) as max_played_at").
		Where("user_id = ?", userId).
		Group("song_id")

	err = p.DB.
		Table("plays AS p").
		Joins("JOIN (?) AS sub ON p.song_id = sub.song_id AND p.played_at = sub.max_played_at", subquery).
		Where("p.user_id = ?", userId).
		Order("p.played_at DESC").
		Preload("Song").
		Preload("Song.Artist").
		Preload("Song.Album").
		Find(&res).Error
	return
}
