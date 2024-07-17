package model

type Career struct {
	Id         int    `gorm:"primaryKey"`
	Title      string `gorm:"not null"`
	Department string `gorm:"not null"`
	Location   string `gorm:"not null"`
}
