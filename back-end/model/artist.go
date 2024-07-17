package model

import (
	"back-end/data/response"
	"time"
)

type Artist struct {
	ArtistId    string  `gorm:"primaryKey"`
	UserId      string  `gorm:"not null"`
	Description *string `gorm:""`
	Banner      *string `gorm:""`
	VerifiedAt  *time.Time
	User        response.UserResponse `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
