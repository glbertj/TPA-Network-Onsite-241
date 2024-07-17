package model

import "back-end/data/response"

type Playlist struct {
	PlaylistId      string `gorm:"primaryKey"`
	UserId          string
	Title           string
	Description     string
	Image           string
	User            response.UserResponse             `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PlaylistDetails []response.PlaylistDetailResponse `gorm:"foreignKey:PlaylistId;references:PlaylistId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
