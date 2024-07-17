package response

type PlayListResponse struct {
	PlaylistId      string                   `json:"playlistId"`
	UserId          string                   `json:"userId"`
	Title           string                   `json:"title"`
	Description     string                   `json:"description"`
	Image           string                   `json:"image"`
	User            UserResponse             `json:"user" gorm:"foreignKey:UserId;references:UserId"`
	PlaylistDetails []PlaylistDetailResponse `json:"playlistDetails" gorm:"foreignKey:PlaylistId;references:PlaylistId"`
}

func (PlayListResponse) TableName() string {
	return "playlists"
}
