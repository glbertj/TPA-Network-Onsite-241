package repository

import "back-end/model"

type AdvertisementRepository interface {
	GetRandomAdvertisement() (res model.Advertisement, err error)
	GetAdvertisementById(advertisementId string) (res model.Advertisement, err error)
}
