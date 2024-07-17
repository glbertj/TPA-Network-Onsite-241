package database

import (
	"back-end/config"
	"back-end/model"
	"back-end/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) *gorm.DB {

	//dbURL := "postgres://postgres:StevenLie@localhost:8080/pre-tpa"
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s", config.Database.Dialect, config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	gormDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil
	}
	err = gormDB.AutoMigrate(
		&model.NotificationSetting{},
		&model.Artist{},
		&model.User{},
		&model.Follow{},
		&model.Album{},
		&model.Song{},
		&model.Play{},
		&model.PlaylistDetails{},
		&model.Playlist{},
		&model.Advertisement{},
	)
	if err != nil {
		utils.CheckError(err)
	}
	return gormDB
}
