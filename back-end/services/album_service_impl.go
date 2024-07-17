package services

import (
	"back-end/data/request"
	"back-end/data/response"
	"back-end/model"
	"back-end/repository"
	"back-end/sse"
	"back-end/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type AlbumServiceImpl struct {
	FollowRepository repository.FollowRepository
	AlbumRepository  repository.AlbumRepository
	ArtistRepository repository.ArtistRepository
	Validate         *validator.Validate
	h                *sse.NotificationSSE
}

func NewAlbumServiceImpl(FollowRepository repository.FollowRepository, AlbumRepository repository.AlbumRepository, ArtistRepository repository.ArtistRepository, Validate *validator.Validate, h *sse.NotificationSSE) *AlbumServiceImpl {
	return &AlbumServiceImpl{FollowRepository: FollowRepository, AlbumRepository: AlbumRepository, ArtistRepository: ArtistRepository, Validate: Validate, h: h}
}

func (a AlbumServiceImpl) GetAlbumsByArtist(artistId string) (res []response.AlbumResponse, err error) {
	resp, err := a.AlbumRepository.GetAlbumsByArtist(artistId)
	if err != nil {
		return nil, err
	}
	for _, album := range resp {
		res = append(res, response.AlbumResponse{
			AlbumId:  album.AlbumId,
			Title:    album.Title,
			Type:     album.Type,
			Banner:   album.Banner,
			Release:  album.Release,
			Artist:   album.Artist,
			ArtistId: album.ArtistId,
		})
	}
	return res, nil
}

func (a AlbumServiceImpl) GetRandomAlbum() (res []response.AlbumResponse, err error) {
	resp, err := a.AlbumRepository.GetRandomAlbum()
	if err != nil {
		return nil, err
	}
	for _, album := range resp {
		res = append(res, response.AlbumResponse{
			AlbumId:  album.AlbumId,
			Title:    album.Title,
			Type:     album.Type,
			Banner:   album.Banner,
			Release:  album.Release,
			Artist:   album.Artist,
			ArtistId: album.ArtistId,
		})
	}
	return res, nil
}

func (a AlbumServiceImpl) CreateAlbum(album request.AlbumRequest) (albumResponse response.AlbumResponse, err error) {
	albumResponse = response.AlbumResponse{
		AlbumId:  utils.GenerateUUID(),
		Title:    album.Title,
		Type:     album.Type,
		Banner:   album.Banner,
		Release:  time.Now(),
		ArtistId: album.ArtistId,
	}

	albumModel := model.Album{
		Title:    albumResponse.Title,
		Type:     albumResponse.Type,
		Banner:   albumResponse.Banner,
		ArtistId: albumResponse.ArtistId,
		Release:  albumResponse.Release,
		AlbumId:  albumResponse.AlbumId,
	}
	err = a.AlbumRepository.CreateAlbum(albumModel)
	if err != nil {
		return
	}
	user, err := a.ArtistRepository.GetArtistByArtistId(albumResponse.ArtistId, true)
	if err != nil {
		return
	}

	fmt.Println("user" + user.UserId)

	follow, err := a.FollowRepository.GetFollower(user.UserId)

	if err != nil {
		return
	}

	for _, v := range follow {
		if v.Follower.NotificationSetting.WebAlbum == true {
			fmt.Println("Create Album " + v.Follower.Email)
			if _, exists := a.h.NotificationChannel[v.FollowerId]; !exists {
				a.h.NotificationChannel[v.FollowerId] = make(chan model.Notification)
			}
			a.h.NotificationChannel[v.FollowerId] <- model.Notification{
				NotifyId: utils.GenerateUUID(),
				UserId:   v.FollowerId,
				Title:    "A new Album has been released",
				Body:     user.User.Username + " has released a new Album (" + album.Title + ") at " + time.Now().Format("2006-01-02"),
				Status:   "OK",
				ReadAt:   time.Now(),
			}
		}
		if v.Follower.NotificationSetting.EmailAlbum == true {
			fmt.Println("Create Album Email" + v.Follower.Email)
			err = utils.SendEmail(v.Follower.Email, "A new Album has been released", user.User.Username+" has released a new Album ("+album.Title+") at "+time.Now().Format("2006-01-02"))
		}
	}

	return
}
