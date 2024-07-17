package main

import (
	"back-end/config"
	"back-end/controller"
	"back-end/database"
	"back-end/repository"
	"back-end/router"
	"back-end/services"
	"back-end/sse"
	"back-end/utils"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {
	cnf := config.LoadEnv()
	db := database.ConnectDB(cnf)
	redis := database.NewRedis(cnf)
	google := database.NewGoogle(cnf)
	validate := validator.New()

	notificationSettingRepository := repository.NewNotificationSettingRepositoryImpl(db, redis)
	notificationSettingService := services.NewNotificationSettingServiceImpl(notificationSettingRepository, validate)
	notificationSettingController := controller.NewNotificationSettingController(notificationSettingService)

	userRepository := repository.NewUserRepositoryImpl(db, redis)
	userService := services.NewUserServiceImpl(userRepository, notificationSettingRepository, validate)
	userController := controller.NewUserController(userService, google)

	notificationChannel := sse.NewNotificationSSE(userRepository)

	playListRepository := repository.NewPlaylistRepositoryImpl(db, redis)
	playListService := services.NewPlaylistServiceImpl(playListRepository, validate)
	playListController := controller.NewPlaylistController(playListService)

	followRepository := repository.NewFollowRepositoryImpl(db, redis)
	followService := services.NewFollowServiceImpl(followRepository, userRepository, validate, notificationChannel)
	followController := controller.NewFollowController(followService)

	songRepository := repository.NewSongRepositoryImpl(db, redis)
	songService := services.NewSongServiceImpl(songRepository, validate)
	songController := controller.NewSongController(songService)

	artistRepository := repository.NewArtistRepositoryImpl(db, redis)
	artistService := services.NewArtistServiceImpl(artistRepository, userRepository, validate)
	artistController := controller.NewArtistController(artistService)

	albumRepository := repository.NewAlbumRepositoryImpl(db, redis)
	albumService := services.NewAlbumServiceImpl(followRepository, albumRepository, artistRepository, validate, notificationChannel)
	albumController := controller.NewAlbumController(albumService)

	queueRepository := repository.NewQueueRepositoryImpl(redis)
	queueService := services.NewQueueServiceImpl(queueRepository, validate)
	queueController := controller.NewQueueController(queueService)

	playRepository := repository.NewPlayRepositoryImpl(db, redis)
	playService := services.NewPlayServiceImpl(playRepository, validate)
	playController := controller.NewPlayController(playService)

	searchService := services.NewSearchService(songRepository, artistRepository, albumRepository, followRepository)
	searchController := controller.NewSearchController(searchService)

	advertisementRepository := repository.NewAdvertisementRepositoryImpl(db, redis)
	advertisementService := services.NewAdvertisementServiceImpl(advertisementRepository, validate)
	advertisementController := controller.NewAdvertisementController(advertisementService)

	r := router.NewRouter(playListController, userController, notificationChannel, followController, songController, albumController, queueController, playController, artistController, notificationSettingController, searchController, advertisementController)

	server := &http.Server{
		Addr:    cnf.Server.Port,
		Handler: r,
	}

	err := server.ListenAndServe()
	utils.CheckError(err)

}
