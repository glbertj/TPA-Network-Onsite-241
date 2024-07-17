package model

type Advertisement struct {
	AdvertisementId string `gorm:"primaryKey"`
	PublisherName   string
	Image           string
	Link            string
}
