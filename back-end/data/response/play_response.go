package response

type PlayResponse struct {
	PlayId string       `json:"playId"`
	SongId string       `json:"songId"`
	UserId string       `json:"userId"`
	User   UserResponse `json:"user" gorm:"foreignKey:UserId;references:UserId"`
	Song   SongResponse `json:"song" gorm:"foreignKey:SongId;references:SongId"`
}

func (PlayResponse) TableName() string {
	return "plays"
}
