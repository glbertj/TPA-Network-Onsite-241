package response

type AdvertisementResponse struct {
	AdvertisementId string `json:"advertisementId"`
	PublisherName   string `json:"publisherName"`
	Image           string `json:"image"`
	Link            string `json:"link"`
}
