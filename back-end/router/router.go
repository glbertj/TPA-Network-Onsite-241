package router

import (
	"back-end/controller"
	"back-end/middleware"
	"back-end/sse"
	"back-end/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(playlist *controller.PlaylistController, user *controller.UserController, h *sse.NotificationSSE, follow *controller.FollowController, song *controller.SongController, album *controller.AlbumController, queue *controller.QueueController, play *controller.PlayController, artist *controller.ArtistController, setting *controller.NotificationSettingController, search *controller.SearchController, adv *controller.AdvertisementController) *gin.Engine {
	router := gin.Default()
	cnf := config.LoadEnv()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4000", cnf.Google.RedirectURL, cnf.Google.AuthURL},
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
		authGroup.GET("/user", user.GetCurrentUser)
		authGroup.GET("/sse/notification-stream", h.StreamNotification)
		authGroup.GET("/music", song.StreamMusic)
		authGroup.GET("/adv", adv.StreamAdv)
		authGroup.POST("/play/create", play.Create)
		authGroup.POST("/user/edit-prof", user.UpdateUserProfile)
		authGroup.GET("/user/current-user", user.GetCurrentUser)
		authGroup.GET("/user/get", user.GetUserById)
		authGroup.GET("/user/logout", user.Logout)
		authGroup.GET("/album/get-random", album.GetRandomAlbum)
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

	return router
}
