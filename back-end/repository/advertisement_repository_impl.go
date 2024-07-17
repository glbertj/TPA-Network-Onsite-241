package repository

import (
	"back-end/database"
	"back-end/model"
	"gorm.io/gorm"
)

type AdvertisementRepositoryImpl struct {
	DB  *gorm.DB
	rdb *database.Redis
}

func NewAdvertisementRepositoryImpl(DB *gorm.DB, rdb *database.Redis) *AdvertisementRepositoryImpl {
	return &AdvertisementRepositoryImpl{DB: DB, rdb: rdb}
}

func (a AdvertisementRepositoryImpl) GetRandomAdvertisement() (res model.Advertisement, err error) {
	err = a.DB.Order("RANDOM()").Limit(1).Find(&res).Error
	return
}

func (a AdvertisementRepositoryImpl) GetAdvertisementById(advertisementId string) (res model.Advertisement, err error) {
	err = a.DB.Where("advertisement_id = ?", advertisementId).Find(&res).Error
	return
}
