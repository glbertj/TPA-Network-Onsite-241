package services

import "back-end/data/response"

type AdvertisementService interface {
	GetRandomAdvertisement() (res response.AdvertisementResponse, err error)
	GetAdvertisementById(advertisementId string) (res response.AdvertisementResponse, err error)
}
