package response

import "time"

type ArtistResponse struct {
	ArtistId    string       `json:"artistId"`
	UserId      string       `json:"userId"`
	Description *string      `json:"description"`
	Banner      *string      `json:"banner"`
	VerifiedAt  *time.Time   `json:"verifiedAt"`
	User        UserResponse `json:"user" gorm:"foreignKey:UserId;references:UserId"`
}

func (ArtistResponse) TableName() string {
	return "artists"
}
