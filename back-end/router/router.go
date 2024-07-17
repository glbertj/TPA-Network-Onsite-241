package router

import (
	"back-end/controller"
	"back-end/middleware"
	"back-end/model"
	"back-end/sse"
	"back-end/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func NewRouter(playlist *controller.PlaylistController, user *controller.UserController, h *sse.NotificationSSE, follow *controller.FollowController, song *controller.SongController, album *controller.AlbumController, queue *controller.QueueController, play *controller.PlayController, artist *controller.ArtistController, setting *controller.NotificationSettingController, search *controller.SearchController, adv *controller.AdvertisementController) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4000", "http://localhost:5173", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept-Ranges", "Content-Type", "Content-Length", "Authorization", "Origin", "Accept", "Range"},
		AllowCredentials: true,
	}))

	router.Static("/public/images/", "./assets/images")
	router.Static("/public/adv/", "./assets/advertise/image")
	router.Static("/public/songs/", "./assets/songs")

	authGroup := router.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware(user.UserService))
	{
		authGroup.GET("/user", user.GetCurrentUser)                     //v
		authGroup.GET("/sse/notification-stream", h.StreamNotification) //v
		authGroup.GET("/music", song.StreamMusic)                       //v
		authGroup.GET("/adv", adv.StreamAdv)                            //v
		authGroup.POST("/play/create", play.Create)                     //TODO : Check
		authGroup.POST("/user/edit-prof", user.UpdateUserProfile)       //v
		authGroup.GET("/user/current-user", user.GetCurrentUser)        //v
		authGroup.GET("/user/get", user.GetUserById)
		//authGroup.POST("/user/sign-out", user.SignOut)
		authGroup.GET("/user/logout", user.Logout)               //v
		authGroup.GET("/album/get-random", album.GetRandomAlbum) //
		authGroup.GET("/play/get-last", play.Get8LastPlayedSongByUser)
		authGroup.POST("/playlist/create", playlist.CreatePlaylist)
		authGroup.GET("/get-following", follow.GetFollowing)
		authGroup.GET("/user/get-all", user.GetAllUser)
		authGroup.POST("/user/update-pic", user.UpdateProfilePicture)
		authGroup.GET("/playlist", playlist.GetPlaylistByUserId)
		authGroup.GET("/playlist-id", playlist.GetPlaylistById)
		authGroup.POST("/playlist-detail", playlist.CreateDetail)
		authGroup.DELETE("/playlist-detail", playlist.DeletePlaylistDetail)
		authGroup.DELETE("/playlist", playlist.DeletePlaylist)
		authGroup.GET("/get-follower", follow.GetFollower)
		authGroup.GET("/get-mutual", follow.GetMutualFollowing)
		authGroup.PUT("/follow", follow.Create)
		authGroup.DELETE("/follow", follow.DeleteFollow)
		authGroup.GET("/album/get-artist", album.GetAlbumByArtist)
		authGroup.GET("/song/get-all", song.GetAllSong)
		authGroup.GET("/song/get", song.GetSongById)
		authGroup.GET("/song/get-by-artist", song.GetSongByArtist)
		authGroup.GET("/song/get-by-album", song.GetSongByAlbum)
		authGroup.GET("/queue/clear", queue.ClearQueue)
		authGroup.POST("/queue/enqueue", queue.Enqueue)
		authGroup.GET("/queue/dequeue", queue.Dequeue)
		authGroup.GET("/queue/get", queue.GetQueue)
		authGroup.GET("/queue/get-all", queue.GetAllQueue)
		authGroup.POST("/queue/remove", queue.RemoveFromQueue)
		authGroup.GET("/play/get-last-rec", play.GetLastPlayedSongByUser)
		authGroup.GET("/artist/get", artist.GetArtistByUserId)
		authGroup.GET("/artist/get-id", artist.GetArtistByArtistId)
		authGroup.POST("/artist/create", artist.CreateArtist)

		authGroup.GET("/artist/get-unverified", artist.GetUnverifiedArtist)
		authGroup.POST("/setting/update", setting.UpdateSetting)
		authGroup.GET("/search/get", search.Search)
		authGroup.GET("/adv/get", adv.GetRandomAdvertisement)
	}

	artistGroup := router.Group("/artist")
	artistGroup.Use(middleware.RoleMiddleware(user.UserService, "Artist"))
	{
		artistGroup.POST("/album/create", album.CreateAlbum)
		artistGroup.POST("/song/create", song.CreateSong)

	}

	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.RoleMiddleware(user.UserService, "Admin"))
	{
		adminGroup.PUT("/artist/update", artist.UpdateVerifyArtist)
		adminGroup.DELETE("/artist/delete", artist.DeleteArtist)

	}

	router.POST("/user/login", user.Authenticate)
	router.POST("/user/update-ver", user.UpdateVerificationStatus)
	router.GET("/auth/google/callback", user.GoogleCallback)
	router.PUT("/user/register", user.Register)
	router.POST("/user/forgot-password", user.Forgot)
	router.POST("/user/reset-password", user.ResetPassword)
	router.GET("/user/valid-verify", user.GetUserByVerifyLink)
	//router.POST("/user/edit-prof", user.UpdateUserProfile)
	//router.GET("/user/current-user", user.GetCurrentUser)
	//router.GET("/user/get", user.GetUserById)
	//router.GET("/user/get-all", user.GetAllUser)

	//router.GET("/user/logout", user.Logout)
	//router.POST("/user/update-pic", user.UpdateProfilePicture)

	//router.GET("/playlist", playlist.GetPlaylistByUserId)
	//router.GET("/playlist-id", playlist.GetPlaylistById)
	//router.POST("/playlist-detail", playlist.CreateDetail)
	//router.DELETE("/playlist-detail", playlist.DeletePlaylistDetail)
	//router.DELETE("/playlist", playlist.DeletePlaylist)
	//router.POST("/playlist/create", playlist.CreatePlaylist)

	//router.GET("/get-following", follow.GetFollowing)
	//router.GET("/get-follower", follow.GetFollower)
	//router.GET("/get-mutual", follow.GetMutualFollowing)
	//router.PUT("/follow", follow.Create)
	//router.DELETE("/follow", follow.DeleteFollow)

	//router.GET("/album/get-artist", album.GetAlbumByArtist)
	//router.POST("/album/create", album.CreateAlbum)
	//router.GET("/album/get-random", album.GetRandomAlbum)

	//router.POST("/song/create", song.CreateSong)
	//router.GET("/song/get-all", song.GetAllSong)
	//router.GET("/song/get", song.GetSongById)
	//router.GET("/song/get-by-artist", song.GetSongByArtist)
	//router.GET("/song/get-by-album", song.GetSongByAlbum)

	//router.GET("/queue/clear", queue.ClearQueue)
	//router.POST("/queue/enqueue", queue.Enqueue)
	//router.GET("/queue/dequeue", queue.Dequeue)
	//router.GET("/queue/get", queue.GetQueue)
	//router.GET("/queue/get-all", queue.GetAllQueue)
	//router.POST("/queue/remove", queue.RemoveFromQueue)

	//router.GET("/play/get-last", play.Get8LastPlayedSongByUser)
	//router.GET("/play/get-last-rec", play.GetLastPlayedSongByUser)
	//router.POST("/play/create", play.Create)

	//router.GET("/artist/get", artist.GetArtistByUserId)
	//router.GET("/artist/get-id", artist.GetArtistByArtistId)
	//router.POST("/artist/create", artist.CreateArtist)
	//router.PUT("/artist/update", artist.UpdateVerifyArtist)
	//router.DELETE("/artist/delete", artist.DeleteArtist)
	//router.GET("/artist/get-unverified", artist.GetUnverifiedArtist)
	//
	//router.POST("/setting/update", setting.UpdateSetting)
	//
	//router.GET("/search/get", search.Search)
	//
	//router.GET("/adv/get", adv.GetRandomAdvertisement)

	//router.GET("/music", song.StreamMusic)
	//router.GET("/adv", adv.StreamAdv)

	//careerGroup := router.Group("/career")
	//careerGroup.Use(middleware.RoleMiddleware(user.UserService, "JLA"))
	//{
	//	careerGroup.POST("/", career.Create)
	//	careerGroup.GET("/", career.FindAll)
	//}

	//router.POST("/career", middleware.AuthMiddleware(user.UserService), career.Create)
	//router.GET("/career", middleware.RoleMiddleware(user.UserService, "JLA"), career.FindAll)
	//router.GET("sse/notification-stream", h.StreamNotification)

	//TESTING
	router.GET("/send", func(c *gin.Context) {
		id := c.Query("id")
		h.NotificationChannel[id] <- model.Notification{
			NotifyId: utils.GenerateUUID(),
			UserId:   "d3cc72e7-f998-4f2f-b45a-1ab86b8bd233",
			Title:    "Tes",
			Body:     "Pong",
			Status:   "OK",
			ReadAt:   time.Time{},
		}
	})

	//router.GET("/test", song.TestMusic)

	return router
}
