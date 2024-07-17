package model

import "back-end/data/response"

type Play struct {
	PlayId   string `gorm:"primaryKey"`
	SongId   string
	UserId   string
	PlayedAt string
	Song     response.SongResponse `gorm:"foreignKey:SongId;references:SongId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User     response.UserResponse `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
