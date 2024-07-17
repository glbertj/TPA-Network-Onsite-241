package services

import (
	"back-end/data/response"
	"back-end/repository"
	"github.com/go-playground/validator/v10"
)

type AdvertisementServiceImpl struct {
	AdvertisementRepository repository.AdvertisementRepository
	Validate                *validator.Validate
}

func NewAdvertisementServiceImpl(AdvertisementRepository repository.AdvertisementRepository, Validate *validator.Validate) *AdvertisementServiceImpl {
	return &AdvertisementServiceImpl{AdvertisementRepository: AdvertisementRepository, Validate: Validate}
}

func (a AdvertisementServiceImpl) GetRandomAdvertisement() (res response.AdvertisementResponse, err error) {
	resp, err := a.AdvertisementRepository.GetRandomAdvertisement()
	if err != nil {
		return
	}
	res = response.AdvertisementResponse{
		AdvertisementId: resp.AdvertisementId,
		PublisherName:   resp.PublisherName,
		Image:           resp.Image,
		Link:            resp.Link,
	}
	return
}

func (a AdvertisementServiceImpl) GetAdvertisementById(advertisementId string) (res response.AdvertisementResponse, err error) {
	resp, err := a.AdvertisementRepository.GetAdvertisementById(advertisementId)
	if err != nil {
		return
	}
	res = response.AdvertisementResponse{
		AdvertisementId: resp.AdvertisementId,
		PublisherName:   resp.PublisherName,
		Image:           resp.Image,
		Link:            resp.Link,
	}
	return
}
